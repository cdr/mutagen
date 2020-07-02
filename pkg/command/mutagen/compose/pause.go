package compose

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/pkg/command/mutagen/daemon"
	"github.com/mutagen-io/mutagen/pkg/command/mutagen/forward"
	"github.com/mutagen-io/mutagen/pkg/command/mutagen/sync"

	"github.com/mutagen-io/mutagen/pkg/compose"
)

// pauseSessions handles Mutagen session pausing for the project.
func pauseSessions(project *compose.Project) error {
	// Connect to the Mutagen daemon and defer closure of the connection.
	daemonConnection, err := daemon.Connect(true, true)
	if err != nil {
		return fmt.Errorf("unable to connect to Mutagen daemon: %w", err)
	}
	defer daemonConnection.Close()

	// Create a session selection for the project.
	projectSelection := project.SessionSelection()

	// Perform forwarding session pausing.
	fmt.Println("Pausing forwarding sessions")
	if err := forward.PauseWithSelection(daemonConnection, projectSelection); err != nil {
		return fmt.Errorf("forwarding pausing failed: %w", err)
	}

	// Perform synchronization session pausing.
	fmt.Println("Pausing synchronization sessions")
	if err := sync.PauseWithSelection(daemonConnection, projectSelection); err != nil {
		return fmt.Errorf("synchronization pausing failed: %w", err)
	}

	// Success.
	return nil
}

// pauseMain is the entry point for the pause command.
func pauseMain(command *cobra.Command, arguments []string) error {
	// Forbid direct control over the Mutagen service.
	for _, argument := range arguments {
		if argument == compose.MutagenServiceName {
			return errors.New("the Mutagen service should not be controlled directly")
		}
	}

	// Load project metadata and defer the release of project resources.
	project, err := compose.LoadProject(
		composeConfiguration.ProjectFlags,
		composeConfiguration.DaemonConnectionFlags,
	)
	if err != nil {
		return fmt.Errorf("unable to load project: %w", err)
	}
	defer project.Dispose()

	// If no services have been explicitly specified, then we're going to pause
	// the entire project (including the Mutagen service), so pause sessions.
	if len(arguments) == 0 {
		if err := pauseSessions(project); err != nil {
			return fmt.Errorf("unable to pause Mutagen sessions: %w", err)
		}
	}

	// Compute the effective top-level flags that we'll use. We reconstitute
	// flags from the root command, but filter project-related flags and replace
	// them with the fully resolved flags from the loaded project.
	topLevelFlags := reconstituteFlags(ComposeCommand.Flags(), topLevelProjectFlagNames)
	topLevelFlags = append(topLevelFlags, project.TopLevelFlags()...)

	// Compute flags and arguments for the command itself.
	pauseArguments := reconstituteFlags(command.Flags(), nil)
	pauseArguments = append(pauseArguments, arguments...)

	// Perform the pass-through operation.
	return invoke(topLevelFlags, "pause", pauseArguments)
}

// pauseCommand is the pause command.
var pauseCommand = &cobra.Command{
	Use:          "pause",
	RunE:         wrapper(pauseMain),
	SilenceUsage: true,
}

// pauseConfiguration stores configuration for the pause command.
var pauseConfiguration struct {
	// help indicates the presence of the -h/--help flag.
	help bool
}

func init() {
	// Avoid Cobra's built-in help functionality that's triggered when the
	// -h/--help flag is present. We still explicitly register a -h/--help flag
	// below for shell completion support.
	pauseCommand.SetHelpFunc(commandHelp)

	// Grab a handle for the command line flags.
	flags := pauseCommand.Flags()

	// Wire up pause command flags.
	flags.BoolVarP(&pauseConfiguration.help, "help", "h", false, "")
}
