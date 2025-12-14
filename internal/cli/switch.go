package cli

import (
	"fmt"

	"github.com/htmluz/worklog/internal/service"
	"github.com/htmluz/worklog/internal/storage"
	"github.com/spf13/cobra"
)

func newSwitchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "switch <from_task_id> <from_window_id> <to_task_id> <to_task_id>",
		Short: "finaliza uma window trackeada",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			windowID := args[0]
			taskID := args[1]

			// TODO
			store, err := storage.NewJSONStorage()
			if err != nil {
				return fmt.Errorf("erro inicializando o storage %w", err)
			}

			ws := service.NewWorklogService(store)
			if err := ws.Stop(taskID, windowID); err != nil {
				return fmt.Errorf("Erro fechando a task %w", err)
			}

			return nil
		},
	}
	return cmd
}
