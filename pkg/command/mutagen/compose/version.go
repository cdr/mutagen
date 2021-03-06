package compose

import (
	"github.com/spf13/cobra"
)

// versionCommand is the version command.
var versionCommand = &cobra.Command{
	Use:                "version",
	Run:                passthrough,
	SilenceUsage:       true,
	DisableFlagParsing: true,
}

// versionConfiguration stores configuration for the version command.
var versionConfiguration struct {
	// help indicates the presence of the -h/--help flag.
	help bool
	// short indicates the presence of the --short flag.
	short bool
}

func init() {
	// We don't set an explicit help function since we disable flag parsing for
	// this command and simply pass arguments directly through to the underlying
	// command. We still explicitly register a -h/--help flag below for shell
	// completion support.

	// Grab a handle for the command line flags.
	flags := versionCommand.Flags()

	// Wire up flags. We don't bother specifying usage information since we'll
	// shell out to Docker Compose if we need to display help information. In
	// the case of this command, we also disable flag parsing and shell out
	// directly, so we only register these flags to support shell completion.
	flags.BoolVarP(&versionConfiguration.help, "help", "h", false, "")
	flags.BoolVar(&versionConfiguration.short, "short", false, "")
}
