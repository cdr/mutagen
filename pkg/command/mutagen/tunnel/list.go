package tunnel

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"

	"github.com/fatih/color"

	"google.golang.org/grpc"

	"github.com/mutagen-io/mutagen/pkg/command"
	"github.com/mutagen-io/mutagen/pkg/command/mutagen/daemon"

	"github.com/mutagen-io/mutagen/pkg/grpcutil"
	"github.com/mutagen-io/mutagen/pkg/selection"
	tunnelingsvc "github.com/mutagen-io/mutagen/pkg/service/tunneling"
	"github.com/mutagen-io/mutagen/pkg/tunneling"
)

// printTunnelStatus prints the status of a tunnel.
func printTunnelStatus(state *tunneling.State) {
	// Print status.
	statusString := state.Status.Description()
	if state.Tunnel.Paused {
		statusString = color.YellowString("[Paused]")
	}
	fmt.Fprintln(color.Output, "Status:", statusString)

	// Print the last error, if any.
	if state.LastError != "" {
		color.Red("Last error: %s\n", state.LastError)
	}
}

// listWithSelection performs a list operation using the provided daemon
// connection and session selection and then prints status information.
func listWithSelection(
	daemonConnection *grpc.ClientConn,
	selection *selection.Selection,
) error {
	// Perform the list operation.
	tunnelingService := tunnelingsvc.NewTunnelingClient(daemonConnection)
	request := &tunnelingsvc.ListRequest{
		Selection: selection,
	}
	response, err := tunnelingService.List(context.Background(), request)
	if err != nil {
		return grpcutil.PeelAwayRPCErrorLayer(err)
	} else if err = response.EnsureValid(); err != nil {
		return errors.Wrap(err, "invalid list response received")
	}

	// Handle output based on whether or not any tunnels were returned.
	if len(response.TunnelStates) > 0 {
		for _, state := range response.TunnelStates {
			fmt.Println(command.DelimiterLine)
			printTunnel(state, listConfiguration.long)
			printTunnelStatus(state)
		}
		fmt.Println(command.DelimiterLine)
	} else {
		fmt.Println(command.DelimiterLine)
		fmt.Println("No tunnels found")
		fmt.Println(command.DelimiterLine)
	}

	// Success.
	return nil
}

// listMain is the entry point for the list command.
func listMain(_ *cobra.Command, arguments []string) error {
	// Create tunnel selection specification.
	selection := &selection.Selection{
		All:            len(arguments) == 0 && listConfiguration.labelSelector == "",
		Specifications: arguments,
		LabelSelector:  listConfiguration.labelSelector,
	}
	if err := selection.EnsureValid(); err != nil {
		return errors.Wrap(err, "invalid tunnel selection specification")
	}

	// Connect to the daemon and defer closure of the connection.
	daemonConnection, err := daemon.Connect(true, true)
	if err != nil {
		return errors.Wrap(err, "unable to connect to daemon")
	}
	defer daemonConnection.Close()

	// Perform the list operation and print status information.
	return listWithSelection(daemonConnection, selection)
}

// listCommand is the list command.
var listCommand = &cobra.Command{
	Use:          "list [<tunnel>...]",
	Short:        "List existing tunnels and their statuses",
	RunE:         listMain,
	SilenceUsage: true,
}

// listConfiguration stores configuration for the list command.
var listConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// long indicates whether or not to use long-format listing.
	long bool
	// labelSelector encodes a label selector to be used in identifying which
	// tunnels should be paused.
	labelSelector string
}

func init() {
	// Grab a handle for the command line flags.
	flags := listCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&listConfiguration.help, "help", "h", false, "Show help information")

	// Wire up list flags.
	flags.BoolVarP(&listConfiguration.long, "long", "l", false, "Show detailed tunnel information")
	flags.StringVar(&listConfiguration.labelSelector, "label-selector", "", "List tunnels matching the specified label selector")
}
