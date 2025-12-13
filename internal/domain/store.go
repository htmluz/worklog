package domain

// https://open.spotify.com/track/0CBB4ZgkW6qQ9Cesr9VO6r?si=c207f32e8d644beb

import "time"

type Store struct {
	Tasks   map[string]*Task   `json:"tasks"`
	Windows map[string]*Window `json:"windows"`
}

func NewStore() *Store {
	return &Store{
		Tasks:   make(map[string]*Task),
		Windows: make(map[string]*Window),
	}
}

func (s *Store) GetTaskDuration(taskID string) time.Duration {
	task, exists := s.Tasks[taskID]
	if !exists {
		return 0
	}

	var total time.Duration
	for _, winID := range task.WindowIDs {
		if win, ok := s.Windows[winID]; ok {
			total += win.Duration()
		}
	}
	return total
}
