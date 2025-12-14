package domain

// https://open.spotify.com/track/23QyE9GQpXsX9WgEDADMa6?si=ee235d6879104a99

import (
	"time"
)

type Window struct {
	ID        string     `json:"id"`
	Issue     string     `json:"issue"`
	Intervals []Interval `json:"intervals"`
	Closed    bool       `json:"closed"`
}

func (w *Window) isActive() bool {
	if len(w.Intervals) == 0 {
		return false
	}
	return w.Intervals[len(w.Intervals)-1].isActive()
}

func (w *Window) Duration() time.Duration {
	var total time.Duration
	for _, interval := range w.Intervals {
		total += interval.Duration()
	}
	return total
}

func (w *Window) Pause() {
	if !w.isActive() {
		return
	}
	now := time.Now()
	w.Intervals[len(w.Intervals)-1].End = &now
}

func (w *Window) Resume() {
	w.Intervals = append(w.Intervals, Interval{Start: time.Now()})
}

func (w *Window) Close() {
	w.Pause()
	w.Closed = true
}

func (w *Window) Start() {
	if len(w.Intervals) > 0 {
		return
	}
	w.Intervals = append(w.Intervals, Interval{Start: time.Now()})
}
