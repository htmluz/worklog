package domain

// https://open.spotify.com/track/29EdNlJQqStWhNkSGpkuFQ?si=f4813bb7ed8d4e66

import "time"

type Task struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Issue     string     `json:"issue,omitempty"`
	WindowIDs []string   `json:"window_ids"`
	CreatedAt time.Time  `json:"created_at"`
	ClosedAt  *time.Time `json:"closed_at,omitempty"`
}

func (t *Task) isActive() bool {
	return t.ClosedAt == nil
}

func (t *Task) AddWindow(windowID string) {
	t.WindowIDs = append(t.WindowIDs, windowID)
}

func (t *Task) RemoveWindow(windowID string) bool {
	for i, d := range t.WindowIDs {
		if d == windowID {
			t.WindowIDs = append(t.WindowIDs[:i], t.WindowIDs[i+1:]...)
			break
		}
	}
	// retorna true p nois fecha se foi a ultima window tmj
	return len(t.WindowIDs) == 0
}

func (t *Task) Close() {
	now := time.Now()
	t.ClosedAt = &now
}
