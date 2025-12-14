package cli

import (
	"fmt"

	"github.com/htmluz/worklog/internal/service"
	"github.com/htmluz/worklog/internal/storage"
	"github.com/spf13/cobra"
)

func newResumeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resume <window_id> <task_id>",
		Short: "finaliza uma window trackeada",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			windowID := args[0]
			taskID := args[1]

			store, err := storage.NewJSONStorage()
			if err != nil {
				return fmt.Errorf("erro inicializando o storage %w", err)
			}

			ws := service.NewWorklogService(store)
			if err := ws.Resume(taskID, windowID); err != nil {
				return fmt.Errorf("Erro fechando a task %w", err)
			}

			return nil
		},
	}
	return cmd
}
