package monitor

import (
	"fmt"
	"sync"
	"time"
)

type MonitorOption func(m *Monitor) *Monitor

func New(mos ...MonitorOption) *Monitor {
	m := &Monitor{}
	for _, mo := range mos {
		m = mo(m)
	}
	return m
}

type Monitor struct {
	ID            string
	Url           string
	Query         map[string]string
	Method        string
	Interval      time.Duration
	locker        sync.Mutex
	MailBox       chan string
	LastCheckedAt time.Time
}

func (m *Monitor) Check() {
	fmt.Println("check")
	if !m.locker.TryLock() {
		return
	}
	m.LastCheckedAt = time.Now()
	m.locker.Unlock()
}

func (m *Monitor) Start() {
	go func() {
		for {
			fmt.Println("for")
			select {
			case <-time.After(m.Interval):
				m.Check()
			case s := <-m.MailBox:
				fmt.Println(s)
				return
			}
		}
	}()
}

func (m *Monitor) Metics() {
	fmt.Println("metrics")
}

func (m *Monitor) Stop() {
	m.MailBox <- "stop"
}
