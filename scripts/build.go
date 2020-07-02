package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/spf13/pflag"

	"github.com/mutagen-io/mutagen/pkg/command"

	"github.com/mutagen-io/mutagen/pkg/agent"
	"github.com/mutagen-io/mutagen/pkg/mutagen"
)

const (
	// agentPackage is the Go package URL to use for building Mutagen agent
	// binaries.
	agentPackage = "github.com/mutagen-io/mutagen/cmd/mutagen-agent"
	// cliPackage is the Go package URL to use for building Mutagen binaries.
	cliPackage = "github.com/mutagen-io/mutagen/cmd/mutagen"

	// agentBuildSubdirectoryName is the name of the build subdirectory where
	// agent binaries are built.
	agentBuildSubdirectoryName = "agent"
	// cliBuildSubdirectoryName is the name of the build subdirectory where CLI
	// binaries are built.
	cliBuildSubdirectoryName = "cli"
	// releaseBuildSubdirectoryName is the name of the build subdirectory where
	// release bundles are built.
	releaseBuildSubdirectoryName = "release"

	// agentBaseName is the name of the Mutagen agent binary without any path or
	// extension.
	agentBaseName = "mutagen-agent"
	// cliBaseName is the name of the Mutagen binary without any path or
	// extension.
	cliBaseName = "mutagen"

	// minimumMacOSVersion is the minimum version of macOS that we'll support
	// (currently pegged to the oldest version of macOS that Go supports).
	minimumMacOSVersion = "10.11"

	// minimumARMSupport is the value to pass to the GOARM environment variable
	// when building binaries. We currently specify support for ARMv5. This will
	// enable software-based floating point. For our use case, this is totally
	// fine, because we don't have any numeric code, and the resulting binary
	// bloat is very minimal. This won't apply for arm64, which always has
	// hardware-based floating point support. For more information, see:
	// https://github.com/golang/go/wiki/GoArm.
	minimumARMSupport = "5"
)

// Target specifies a GOOS/GOARCH combination.
type Target struct {
	// GOOS is the GOOS environment variable specification for the target.
	GOOS string
	// GOARCH is the GOARCH environment variable specification for the target.
	GOARCH string
}

// String generates a human-readable representation of the target.
func (t Target) String() string {
	return fmt.Sprintf("%s/%s", t.GOOS, t.GOARCH)
}

// Name generates a representation of the target that is suitable for paths and
// file names.
func (t Target) Name() string {
	return fmt.Sprintf("%s_%s", t.GOOS, t.GOARCH)
}

// ExecutableName formats executable names for the target.
func (t Target) ExecutableName(base string) string {
	// If we're on Windows, append a ".exe" extension.
	if t.GOOS == "windows" {
		return fmt.Sprintf("%s.exe", base)
	}

	// Otherwise return the base name unmodified.
	return base
}

// goEnv generates an environment that can be used when invoking the Go
// toolchain to generate output for the target.
func (t Target) goEnv() ([]string, error) {
	// Duplicate the existing environment.
	result := os.Environ()

	// Force use of Go modules.
	result = append(result, "GO111MODULE=on")

	// Override GOOS/GOARCH.
	result = append(result, fmt.Sprintf("GOOS=%s", t.GOOS))
	result = append(result, fmt.Sprintf("GOARCH=%s", t.GOARCH))

	// If we're on macOS, we're going to use cgo to access the FSEvents API, so
	// we need to ensure that we compile with flags that tell the C compiler and
	// external linker to support older versions of macOS. These flags will tell
	// the C compiler to generate code compatible with the target version of
	// macOS and tell the external linker what value to embed for the
	// LC_VERSION_MIN_MACOSX flag in the resulting Mach-O binaries. Go's
	// internal linker automatically defaults to a relatively liberal (old)
	// value for this flag, but since we're using an external linker, it
	// defaults to the current SDK version.
	//
	// For all other platforms, we disable cgo. This is essential for our Linux
	// CI setup, because we build agent executables during testing that we then
	// run inside Docker containers for our integration tests. These containers
	// typically run Alpine Linux, and if the agent binary is linked to C
	// libraries that only exist on the build system, then they won't work
	// inside the container. We can't disable cgo on a global basis though,
	// because it's needed for race condition testing. Another reason that it's
	// good to disable cgo when building agent binaries during testing is that
	// the release agent binaries will also have cgo disabled (except on macOS),
	// and we'll want to faithfully recreate that.
	//
	// TODO: If we ever decide to build for iOS, we should set the corresponding
	// flags for that platform.
	if t.GOOS == "darwin" && t.GOARCH == "amd64" {
		result = append(result, fmt.Sprintf("CGO_CFLAGS=-mmacosx-version-min=%s", minimumMacOSVersion))
		result = append(result, fmt.Sprintf("CGO_LDFLAGS=-mmacosx-version-min=%s", minimumMacOSVersion))
	} else {
		result = append(result, "CGO_ENABLED=0")
	}

	// Set up ARM target support. See notes for definition of minimumARMSupport.
	// We don't need to unset any existing GOARM variables since they simply
	// won't be used if we're not targeting (non-64-bit) ARM systems.
	if t.GOARCH == "arm" {
		result = append(result, fmt.Sprintf("GOARM=%s", minimumARMSupport))
	}

	// Done.
	return result, nil
}

// IsCrossTarget determines whether or not the target represents a
// cross-compilation target (i.e. not the native target for the current Go
// toolchain).
func (t Target) IsCrossTarget() bool {
	return t.GOOS != runtime.GOOS || t.GOARCH != runtime.GOARCH
}

// IncludeAgentInSlimBuildModes indicates whether or not the target should have
// an agent binary included in the agent bundle in slim and release-slim modes.
func (t Target) IncludeAgentInSlimBuildModes() bool {
	return !t.IsCrossTarget() ||
		(t.GOOS == "darwin" && t.GOARCH == "amd64") ||
		(t.GOOS == "windows" && t.GOARCH == "amd64") ||
		(t.GOOS == "linux" && (t.GOARCH == "amd64" || t.GOARCH == "arm")) ||
		(t.GOOS == "freebsd" && t.GOARCH == "amd64")
}

// BuildBundleInReleaseSlimMode indicates whether or not the target should have
// a release bundle built in release-slim mode.
func (t Target) BuildBundleInReleaseSlimMode() bool {
	return !t.IsCrossTarget() ||
		(t.GOOS == "darwin" && t.GOARCH == "amd64") ||
		(t.GOOS == "windows" && t.GOARCH == "amd64") ||
		(t.GOOS == "linux" && t.GOARCH == "amd64")
}

// IsContainerTarget indicates whether or not a target is used for containers.
// It doesn't mean that the target is used exclusively for containers, but it
// means that individual CLI and agent release bundles should be generated for
// the target (in addition to the standard release bundles).
func (t Target) IsContainerTarget() bool {
	return t.GOOS == "linux" && t.GOARCH == "amd64"
}

// Build executes a module-aware build of the specified package URL, storing the
// output of the build at the specified path.
func (t Target) Build(url, output string) error {
	// Create the build command. We use the "-s -w" linker flags to omit the
	// symbol table and debugging information. This shaves off about 25% of the
	// binary size and only disables debugging (stack traces are still intact).
	// For more information, see:
	// https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick
	builder := exec.Command("go", "build", "-mod=readonly", "-o", output, "-ldflags=-s -w", url)

	// Set the environment.
	environment, err := t.goEnv()
	if err != nil {
		return errors.Wrap(err, "unable to create build environment")
	}
	builder.Env = environment

	// Forward input, output, and error streams.
	builder.Stdin = os.Stdin
	builder.Stdout = os.Stdout
	builder.Stderr = os.Stderr

	// Run the build.
	return errors.Wrap(builder.Run(), "compilation failed")
}

// targets encodes which combinations of GOOS and GOARCH we want to use for
// building agent and CLI binaries. We don't build every target at the moment,
// but we do list all potential targets here and comment out those we don't
// support. This list is created from https://golang.org/doc/install/source.
// Unfortunately there's no automated way to construct this list, but that's
// fine since we have to manually groom it anyway.
var targets = []Target{
	// We completely disable Android because it doesn't provide a useful shell
	// or SSH server.
	// {"android", "arm"},
	// We completely disable darwin/386 because Go only supports macOS 10.7+,
	// which is always able to run amd64 binaries.
	// {"darwin", "386"},
	{"darwin", "amd64"},
	// We completely disble darwin/arm and darwin/arm64 because no ARM-based
	// Darwin platforms (iOS, watchOS, tvOS) provide a useful shell or SSH
	// server.
	// {"darwin", "arm"},
	// TODO: Figure out why darwin/arm64 doesn't compile in any case anyway.
	// We're seeing this issue: https://github.com/golang/go/issues/16445. The
	// "resolution" there makes sense, except that darwin/arm compiles fine.
	// According to https://golang.org/cmd/cgo/, CGO is automatically disabled
	// when cross-compiling. It's not clear if CGO is being disabled for
	// darwin/arm (and not for darwin/arm64) or if the environment is just
	// broken for darwin/arm64. In either case, there's some bug that needs to
	// be fixed. Either CGO needs to be disabled automatically in both cases for
	// consistency, or the cross-compilation environment needs to be fixed.
	// {"darwin", "arm64"},
	{"dragonfly", "amd64"},
	{"freebsd", "386"},
	{"freebsd", "amd64"},
	{"freebsd", "arm"},
	{"linux", "386"},
	{"linux", "amd64"},
	{"linux", "arm"},
	{"linux", "arm64"},
	{"linux", "ppc64"},
	{"linux", "ppc64le"},
	{"linux", "mips"},
	{"linux", "mipsle"},
	{"linux", "mips64"},
	{"linux", "mips64le"},
	// TODO: This combination is valid but not listed on the "Installing Go from
	// source" page. Perhaps we should open a pull request to change that?
	{"linux", "s390x"},
	{"netbsd", "386"},
	{"netbsd", "amd64"},
	{"netbsd", "arm"},
	{"openbsd", "386"},
	{"openbsd", "amd64"},
	{"openbsd", "arm"},
	// We completely disable Plan 9 because it is just missing too many
	// facilities for Mutagen to build easily, even just the agent component.
	// TODO: We might be able to get Plan 9 functioning as an agent, but it's
	// going to take some serious playing around with source file layouts and
	// build tags. To get started looking into this, look for the !plan9 build
	// tag and see where the gaps are. Most of the problems revolve around the
	// syscall package, but none of that is necessary for the agent, so it can
	// probably be built.
	// {"plan9", "386"},
	// {"plan9", "amd64"},
	{"solaris", "amd64"},
	{"windows", "386"},
	{"windows", "amd64"},
	{"windows", "arm"},
}

// archiveBuilderCopyBufferSize determines the size of the copy buffer used when
// generating archive files.
// TODO: Figure out if we should set this on a per-machine basis. This value is
// taken from Go's io.Copy method, which defaults to allocating a 32k buffer if
// none is provided.
const archiveBuilderCopyBufferSize = 32 * 1024

type ArchiveBuilder struct {
	file       *os.File
	compressor *gzip.Writer
	archiver   *tar.Writer
	copyBuffer []byte
}

func NewArchiveBuilder(bundlePath string) (*ArchiveBuilder, error) {
	// Open the underlying file.
	file, err := os.Create(bundlePath)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create target file")
	}

	// Create the compressor.
	compressor, err := gzip.NewWriterLevel(file, gzip.BestCompression)
	if err != nil {
		file.Close()
		return nil, errors.Wrap(err, "unable to create compressor")
	}

	// Success.
	return &ArchiveBuilder{
		file:       file,
		compressor: compressor,
		archiver:   tar.NewWriter(compressor),
		copyBuffer: make([]byte, archiveBuilderCopyBufferSize),
	}, nil
}

func (b *ArchiveBuilder) Close() error {
	// Close in the necessary order to trigger flushes.
	if err := b.archiver.Close(); err != nil {
		b.compressor.Close()
		b.file.Close()
		return errors.Wrap(err, "unable to close archiver")
	} else if err := b.compressor.Close(); err != nil {
		b.file.Close()
		return errors.Wrap(err, "unable to close compressor")
	} else if err := b.file.Close(); err != nil {
		return errors.Wrap(err, "unable to close file")
	}

	// Success.
	return nil
}

func (b *ArchiveBuilder) Add(name, path string, mode int64) error {
	// If the name is empty, use the base name.
	if name == "" {
		name = filepath.Base(path)
	}

	// Open the file and ensure its cleanup.
	file, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "unable to open file")
	}
	defer file.Close()

	// Compute its size.
	stat, err := file.Stat()
	if err != nil {
		return errors.Wrap(err, "unable to determine file size")
	}
	size := stat.Size()

	// Write the header for the entry.
	header := &tar.Header{
		Name:    name,
		Mode:    mode,
		Size:    size,
		ModTime: time.Now(),
	}
	if err := b.archiver.WriteHeader(header); err != nil {
		return errors.Wrap(err, "unable to write archive header")
	}

	// Copy the file contents.
	if _, err := io.CopyBuffer(b.archiver, file, b.copyBuffer); err != nil {
		return errors.Wrap(err, "unable to write archive entry")
	}

	// Success.
	return nil
}

// copyFile copies the contents at sourcePath to a newly created file at
// destinationPath that inherits the permissions of sourcePath.
func copyFile(sourcePath, destinationPath string) error {
	// Open the source file and defer its closure.
	source, err := os.Open(sourcePath)
	if err != nil {
		return errors.Wrap(err, "unable to open source file")
	}
	defer source.Close()

	// Grab source file metadata.
	metadata, err := source.Stat()
	if err != nil {
		return errors.Wrap(err, "unable to query source file metadata")
	}

	// Remove the destination.
	os.Remove(destinationPath)

	// Create the destination file and defer its closure. We open with exclusive
	// creation flags to ensure that we're the ones creating the file so that
	// its permissions are set correctly.
	destination, err := os.OpenFile(destinationPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, metadata.Mode()&os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "unable to create destination file")
	}
	defer destination.Close()

	// Copy contents.
	if count, err := io.Copy(destination, source); err != nil {
		return errors.Wrap(err, "unable to copy data")
	} else if count != metadata.Size() {
		return errors.New("copied size does not match expected")
	}

	// Success.
	return nil
}

var usage = `usage: build [-h|--help] [-m|--mode=<mode>]

The mode flag accepts four values: 'local', 'slim', 'release', and
'release-slim'. 'local' will build CLI and agent binaries only for the current
platform. 'slim' will build the CLI binary for only the current platform and
agents for a small subset of platforms. 'release' will build CLI and agent
binaries for all platforms and package for release. 'release-slim' is the same
as release but only builds release bundles for a small subset of platforms. The
default mode is 'local'.
`

// build is the primary entry point.
func build() error {
	// Parse command line arguments.
	flagSet := pflag.NewFlagSet("build", pflag.ContinueOnError)
	flagSet.SetOutput(ioutil.Discard)
	var mode string
	flagSet.StringVarP(&mode, "mode", "m", "local", "specify the build mode")
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		if err == pflag.ErrHelp {
			fmt.Fprint(os.Stdout, usage)
			return nil
		} else {
			return errors.Wrap(err, "unable to parse command line")
		}
	}
	if !(mode == "local" || mode == "slim" || mode == "release" || mode == "release-slim") {
		return fmt.Errorf("invalid build mode: %s", mode)
	}

	// The only platform really suited to cross-compiling for every other
	// platform at the moment is macOS. This is because its DNS resolution
	// really has to be done through the system's DNS resolver in order to
	// function properly and because FSEvents is used for file monitoring and
	// that is a C-based API, not accessible purely via system calls. All of the
	// other platforms can survive with pure Go compilation.
	if runtime.GOOS != "darwin" {
		if mode == "release" {
			return errors.New("macOS required for release builds")
		} else if mode == "slim" || mode == "release-slim" {
			command.Warning("macOS agents will be built without cgo support")
		}
	}

	// Compute the path to the Mutagen source directory.
	mutagenSourcePath, err := mutagen.SourceTreePath()
	if err != nil {
		return errors.Wrap(err, "unable to compute Mutagen source tree path")
	}

	// Verify that we're running inside the Mutagen source directory, otherwise
	// we can't rely on Go modules working.
	workingDirectory, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "unable to compute working directory")
	}
	workingDirectoryRelativePath, err := filepath.Rel(mutagenSourcePath, workingDirectory)
	if err != nil {
		return errors.Wrap(err, "unable to determine working directory relative path")
	}
	if strings.Contains(workingDirectoryRelativePath, "..") {
		return errors.New("build script run outside Mutagen source tree")
	}

	// Compute the path to the build directory and ensure that it exists.
	buildPath := filepath.Join(mutagenSourcePath, mutagen.BuildDirectoryName)
	if err := os.MkdirAll(buildPath, 0700); err != nil {
		return errors.Wrap(err, "unable to create build directory")
	}

	// Create the necessary build directory hierarchy.
	agentBuildSubdirectoryPath := filepath.Join(buildPath, agentBuildSubdirectoryName)
	cliBuildSubdirectoryPath := filepath.Join(buildPath, cliBuildSubdirectoryName)
	releaseBuildSubdirectoryPath := filepath.Join(buildPath, releaseBuildSubdirectoryName)
	if err := os.MkdirAll(agentBuildSubdirectoryPath, 0700); err != nil {
		return errors.Wrap(err, "unable to create agent build subdirectory")
	}
	if err := os.MkdirAll(cliBuildSubdirectoryPath, 0700); err != nil {
		return errors.Wrap(err, "unable to create CLI build subdirectory")
	}
	if mode == "release" || mode == "release-slim" {
		if err := os.MkdirAll(releaseBuildSubdirectoryPath, 0700); err != nil {
			return errors.Wrap(err, "unable to create release build subdirectory")
		}
	}

	// Compute the local target.
	localTarget := Target{runtime.GOOS, runtime.GOARCH}

	// Compute agent targets.
	var agentTargets []Target
	for _, target := range targets {
		if mode == "local" && target.IsCrossTarget() {
			continue
		} else if (mode == "slim" || mode == "release-slim") && !target.IncludeAgentInSlimBuildModes() {
			continue
		}
		agentTargets = append(agentTargets, target)
	}

	// Compute CLI targets.
	var cliTargets []Target
	for _, target := range targets {
		if (mode == "local" || mode == "slim") && target.IsCrossTarget() {
			continue
		} else if mode == "release-slim" && !target.BuildBundleInReleaseSlimMode() {
			continue
		}
		cliTargets = append(cliTargets, target)
	}

	// Build agent binaries.
	log.Println("Building agent binaries...")
	for _, target := range agentTargets {
		log.Println("Building agent for", target)
		agentBuildPath := filepath.Join(agentBuildSubdirectoryPath, target.Name())
		if err := target.Build(agentPackage, agentBuildPath); err != nil {
			return errors.Wrap(err, "unable to build agent")
		}
	}

	// Build CLI binaries.
	log.Println("Building CLI binaries...")
	for _, target := range cliTargets {
		log.Println("Build CLI for", target)
		cliBuildPath := filepath.Join(cliBuildSubdirectoryPath, target.Name())
		if err := target.Build(cliPackage, cliBuildPath); err != nil {
			return errors.Wrap(err, "unable to build CLI")
		}
	}

	// Build the agent bundle.
	log.Println("Building agent bundle...")
	agentBundlePath := filepath.Join(buildPath, agent.BundleName)
	agentBundleBuilder, err := NewArchiveBuilder(agentBundlePath)
	if err != nil {
		return errors.Wrap(err, "unable to create agent bundle archive builder")
	}
	for _, target := range agentTargets {
		agentBuildPath := filepath.Join(agentBuildSubdirectoryPath, target.Name())
		if err := agentBundleBuilder.Add(target.Name(), agentBuildPath, 0755); err != nil {
			agentBundleBuilder.Close()
			return errors.Wrap(err, "unable to add agent to bundle")
		}
	}
	if err := agentBundleBuilder.Close(); err != nil {
		return errors.Wrap(err, "unable to finalize agent bundle")
	}

	// Build release bundles if necessary.
	if mode == "release" || mode == "release-slim" {
		log.Println("Building release bundles...")
		for _, target := range cliTargets {
			// Update status.
			log.Println("Building release bundle for", target)

			// Compute paths.
			cliBuildPath := filepath.Join(cliBuildSubdirectoryPath, target.Name())
			releaseBundlePath := filepath.Join(
				releaseBuildSubdirectoryPath,
				fmt.Sprintf("mutagen_%s_v%s.tar.gz", target.Name(), mutagen.Version),
			)

			// Build the release bundle.
			if releaseBundle, err := NewArchiveBuilder(releaseBundlePath); err != nil {
				return errors.Wrap(err, "unable to create release bundle")
			} else if err = releaseBundle.Add(target.ExecutableName(cliBaseName), cliBuildPath, 0755); err != nil {
				releaseBundle.Close()
				return errors.Wrap(err, "unable to add CLI to release bundle")
			} else if err = releaseBundle.Add("", agentBundlePath, 0644); err != nil {
				releaseBundle.Close()
				return errors.Wrap(err, "unable to add agent bundle to release bundle")
			} else if err = releaseBundle.Close(); err != nil {
				return errors.Wrap(err, "unable to finalize release bundle")
			}

			// If this is a container platform, then build individual CLI and
			// agent archives.
			if target.IsContainerTarget() {
				// Update status.
				log.Println("Building archives for", target)

				// Compute additional paths.
				agentBuildPath := filepath.Join(agentBuildSubdirectoryPath, target.Name())
				agentArchivePath := filepath.Join(
					releaseBuildSubdirectoryPath,
					fmt.Sprintf("agent_%s_v%s.tar.gz", target.Name(), mutagen.Version),
				)
				cliArchivePath := filepath.Join(
					releaseBuildSubdirectoryPath,
					fmt.Sprintf("cli_%s_v%s.tar.gz", target.Name(), mutagen.Version),
				)

				// Build the agent archive.
				if agentArchive, err := NewArchiveBuilder(agentArchivePath); err != nil {
					return errors.Wrap(err, "unable to create agent archive")
				} else if err = agentArchive.Add(target.ExecutableName(agentBaseName), agentBuildPath, 0755); err != nil {
					agentArchive.Close()
					return errors.Wrap(err, "unable to add agent to agent archive")
				} else if err = agentArchive.Close(); err != nil {
					return errors.Wrap(err, "unable to finalize agent archive")
				}

				// Build the CLI archive.
				if cliArchive, err := NewArchiveBuilder(cliArchivePath); err != nil {
					return errors.Wrap(err, "unable to create CLI archive")
				} else if err = cliArchive.Add(target.ExecutableName(cliBaseName), cliBuildPath, 0755); err != nil {
					cliArchive.Close()
					return errors.Wrap(err, "unable to add CLI to CLI archive")
				} else if err = cliArchive.Close(); err != nil {
					return errors.Wrap(err, "unable to finalize CLI archive")
				}
			}
		}
	}

	// Relocate the CLI binary for the current platform.
	log.Println("Copying binary for testing")
	localCLIBuildPath := filepath.Join(cliBuildSubdirectoryPath, localTarget.Name())
	localCLIRelocationPath := filepath.Join(buildPath, localTarget.ExecutableName(cliBaseName))
	if err := copyFile(localCLIBuildPath, localCLIRelocationPath); err != nil {
		return errors.Wrap(err, "unable to copy current platform CLI")
	}

	// Success.
	return nil
}

func main() {
	if err := build(); err != nil {
		command.Fatal(err)
	}
}
