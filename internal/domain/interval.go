package domain

// https://open.spotify.com/track/23QyE9GQpXsX9WgEDADMa6?si=02be587bf8494054

import "time"

type Interval struct {
	Start time.Time  `json:"start"`
	End   *time.Time `json:"end,omitempty"`
}

func (i Interval) isActive() bool {
	return i.End == nil
}

func (i Interval) Duration() time.Duration {
	end := time.Now()
	if i.End != nil {
		end = *i.End
	}
	return end.Sub(i.Start)
}
