package domain

// https://open.spotify.com/track/29EdNlJQqStWhNkSGpkuFQ?si=f4813bb7ed8d4e66

import "time"

type Task struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Issue         string     `json:"issue,omitempty"`
	Windows       []*Window  `json:"windows"`
	ClosedWindows []*Window  `json:"closed_windows,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	ClosedAt      *time.Time `json:"closed_at,omitempty"`
}

func (t *Task) isActive() bool {
	return t.ClosedAt == nil
}

func (t *Task) AddWindow(w Window) {
	t.Windows = append(t.Windows, &w)
}

func (t *Task) RemoveWindow(windowID string) bool {
	for i, w := range t.Windows {
		if w.ID == windowID {
			w.Close()

			t.ClosedWindows = append(t.ClosedWindows, w)
			t.Windows = append(t.Windows[:i], t.Windows[i+1:]...)
			break
		}
	}
	// retorna true p nois fecha se foi a ultima window tmj
	return len(t.Windows) == 0
}

func (t *Task) Close() {
	now := time.Now()
	t.ClosedAt = &now

	for i := range t.Windows {
		t.Windows[i].Close()
		t.ClosedWindows = append(t.ClosedWindows, t.Windows[i])
	}
	t.Windows = nil
}
