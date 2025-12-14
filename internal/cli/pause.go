package cli

import (
	"fmt"

	"github.com/htmluz/worklog/internal/service"
	"github.com/htmluz/worklog/internal/storage"
	"github.com/spf13/cobra"
)

func newPauseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "pause <window_id> <task_id>",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			windowID := args[0]
			taskID := args[1]

			store, err := storage.NewJSONStorage()
			if err != nil {
				return fmt.Errorf("erro inicializando o storage %w", err)
			}

			ws := service.NewWorklogService(store)
			if err := ws.Pause(taskID, windowID); err != nil {
				return fmt.Errorf("Erro pausando a task %w", err)
			}

			return nil
		},
	}
	return cmd
}
