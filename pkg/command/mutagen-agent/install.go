package main

import (
	"github.com/pkg/errors"

	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/pkg/command"

	"github.com/mutagen-io/mutagen/pkg/agent"
)

// installMain is the entry point for the install command.
func installMain(_ *cobra.Command, _ []string) error {
	return errors.Wrap(agent.Install(), "installation error")
}

// installCommand is the install command.
var installCommand = &cobra.Command{
	Use:          agent.ModeInstall,
	Short:        "Perform agent installation",
	Args:         command.DisallowArguments,
	RunE:         installMain,
	SilenceUsage: true,
}

// installConfiguration stores configuration for the install command.
var installConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
}

func init() {
	// Grab a handle for the command line flags.
	flags := installCommand.Flags()

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&installConfiguration.help, "help", "h", false, "Show help information")
}
