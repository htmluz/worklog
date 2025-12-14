package cli

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "worklog",
		Short: ":3",
	}
	cmd.AddCommand(newStartCmd())
	return cmd
}
