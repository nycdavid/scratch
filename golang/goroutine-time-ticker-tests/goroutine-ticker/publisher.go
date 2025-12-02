package goroutine_ticker

import (
	"fmt"
	"time"
)

type (
	Printer struct {
		Iteration int
	}

	Ticker interface {
		C() <-chan time.Time
		Stop()
	}

	realClock struct {
		t *time.Ticker
	}
)

func (r *realClock) C() <-chan time.Time {
	return r.t.C
}

func (r *realClock) Stop() {
	r.t.Stop()
}

func (p *Printer) Start() {
	t := time.NewTicker(10 * time.Second) // fixed real interval
	defer t.Stop()

	for {
		select {
		case <-t.C:
			// Directly tied to wall clock, hard to control in tests
			fmt.Println("tick at", time.Now())
			p.Iteration++
		}
	}
}
