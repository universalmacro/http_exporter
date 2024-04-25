package monitor

import (
	"testing"
	"time"
)

func TestMonitor(t *testing.T) {
	// This is a placeholder for a test function.
	monitor := New(
		func(m *Monitor) *Monitor {
			m.ID = "1"
			return m
		},
		func(m *Monitor) *Monitor {
			m.Url = "http://localhost:8080"
			return m
		},
		func(m *Monitor) *Monitor {
			m.Interval = 10 * time.Second
			return m
		},
	)
	monitor.Start()
	time.Sleep(21 * time.Second)
	monitor.Stop()
}
