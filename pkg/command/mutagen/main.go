package mutagen

import (
	"os"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/fatih/color"

	"github.com/mutagen-io/mutagen/pkg/command"
	"github.com/mutagen-io/mutagen/pkg/command/mutagen/compose"
	"github.com/mutagen-io/mutagen/pkg/command/mutagen/daemon"
	"github.com/mutagen-io/mutagen/pkg/command/mutagen/forward"
	"github.com/mutagen-io/mutagen/pkg/command/mutagen/project"
	"github.com/mutagen-io/mutagen/pkg/command/mutagen/sync"
	"github.com/mutagen-io/mutagen/pkg/command/mutagen/tunnel"

	"github.com/mutagen-io/mutagen/pkg/prompting"
)

// rootMain is the entry point for the root command.
func rootMain(command *cobra.Command, _ []string) error {
	// If no commands were given, then print help information and bail. We don't
	// have to worry about warning about arguments being present here (which
	// would be incorrect usage) because arguments can't even reach this point
	// (they will be mistaken for subcommands and a error will be displayed).
	command.Help()

	// Success.
	return nil
}

// rootCommand is the root command.
var rootCommand = &cobra.Command{
	Use:          "mutagen",
	Short:        "Mutagen is a remote development tool built on fast file synchronization",
	RunE:         rootMain,
	SilenceUsage: true,
}

// rootConfiguration stores configuration for the root command.
var rootConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
}

func init() {
	// Disable alphabetical sorting of commands in help output. This is a global
	// setting that affects all Cobra command instances.
	cobra.EnableCommandSorting = false

	// Disable Cobra's use of mousetrap. This breaks daemon registration on
	// Windows because it tries to enforce that the CLI only be launched from
	// a console, which it's not when running automatically.
	cobra.MousetrapHelpText = ""

	// Grab a handle for the command line flags.
	flags := rootCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&rootConfiguration.help, "help", "h", false, "Show help information")

	// Register commands.
	// HACK: Add the sync commands as direct subcommands of the root command for
	// temporary backward compatibility.
	commands := []*cobra.Command{
		sync.SyncCommand,
		forward.ForwardCommand,
		project.ProjectCommand,
		compose.RootCommand,
		tunnel.TunnelCommand,
		loginCommand,
		logoutCommand,
		daemon.DaemonCommand,
		versionCommand,
		legalCommand,
		generateCommand,
	}
	commands = append(commands, sync.Commands...)
	rootCommand.AddCommand(commands...)

	// HACK: Register the sync subcommands with the sync command after
	// registering them with the root command so that they have the correct
	// parent command and thus the correct help output.
	sync.SyncCommand.AddCommand(sync.Commands...)

	// HACK If we're on Windows, enable color support for command usage and
	// error output by recursively replacing the output streams for Cobra
	// commands.
	if runtime.GOOS == "windows" {
		enableColorForCommand(rootCommand)
	}
}

// enableColorForCommand recursively enables colorized usage and error output
// for a command and all of its child commands.
func enableColorForCommand(command *cobra.Command) {
	// Enable color support for the command itself.
	command.SetOut(color.Output)
	command.SetErr(color.Error)

	// Recursively enable color support for child commands.
	for _, c := range command.Commands() {
		enableColorForCommand(c)
	}
}

func Main() {
	// Check if a prompting environment is set. If so, treat this as a prompt
	// request. Prompting is sort of a special pseudo-command that's indicated
	// by the presence of an environment variable, and hence it has to be
	// handled in a bit of a special manner.
	if _, ok := os.LookupEnv(prompting.PrompterEnvironmentVariable); ok {
		if err := promptMain(os.Args[1:]); err != nil {
			command.Fatal(err)
		}
		return
	}

	// Handle terminal compatibility issues. If this call returns, it means that
	// we should proceed normally.
	command.HandleTerminalCompatibility()

	// HACK: If we're performing command line completion, then remove the
	// adapter command that we use to keep the Docker Compose command hierarchy
	// separate and replace it with the actual Docker Compose command hierarchy
	// so that completions work properly.
	if command.PerformingShellCompletion {
		rootCommand.RemoveCommand(compose.RootCommand)
		rootCommand.AddCommand(compose.ComposeCommand)
	}

	// HACK: Modify the root command help to hide legacy root sync commands.
	defaultHelpFunction := rootCommand.HelpFunc()
	rootCommand.SetHelpFunc(func(command *cobra.Command, arguments []string) {
		if command == rootCommand {
			for _, command := range sync.Commands {
				command.Hidden = true
			}
		}
		defaultHelpFunction(command, arguments)
	})

	// Execute the root command.
	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
