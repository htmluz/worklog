package cli

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "worklog",
		Short: ":3",
	}
	cmd.AddCommand(newStartCmd())
	cmd.AddCommand(newStopCmd())
	cmd.AddCommand(newChildCmd())
	cmd.AddCommand(newPauseCmd())
	cmd.AddCommand(newResumeCmd())

	return cmd
}
