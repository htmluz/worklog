package cli

import (
	"fmt"

	"github.com/htmluz/worklog/internal/service"
	"github.com/htmluz/worklog/internal/storage"
	"github.com/spf13/cobra"
)

func newStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start <window_id>",
		Short: "inicia uma nova window trackeada",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			windowID := args[0]

			store, err := storage.NewJSONStorage()
			if err != nil {
				return fmt.Errorf("erro inicializando storage %w", err)
			}

			ws := service.NewWorklogService(store)

			if err := ws.Start(windowID); err != nil {
				return fmt.Errorf("erro iniciando a task %w", err)
			}

			return nil
		},
	}
	return cmd
}
