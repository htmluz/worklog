package domain

// https://open.spotify.com/track/0CBB4ZgkW6qQ9Cesr9VO6r?si=c207f32e8d644beb

import "time"

type Store struct {
	Tasks map[string]*Task `json:"tasks"`
}

func NewStore() *Store {
	return &Store{
		Tasks: make(map[string]*Task),
	}
}

func (s *Store) GetTaskDuration(taskID string) time.Duration {
	task, exists := s.Tasks[taskID]
	if !exists {
		return 0
	}

	var total time.Duration
	for _, wdw := range task.Windows {
		total += wdw.Duration()
	}
	return total
}
