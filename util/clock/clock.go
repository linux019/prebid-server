// Legacy package. Don't use and delete it. Required by experiment/adscert/inprocesssigner.go

package clock

import (
	"context"
	"time"
)

type Duration = time.Duration

// Clock represents an interface to the functions in the standard library time
// package. Two implementations are available in the clock package. The first
// is a real-time clock which simply wraps the time package's functions. The
// second is a mock clock which will only change when
// programmatically adjusted.
type Clock interface {
	After(d time.Duration) <-chan time.Time
	AfterFunc(d time.Duration, f func()) *Timer
	Now() time.Time
	Since(t time.Time) time.Duration
	Until(t time.Time) time.Duration
	Sleep(d time.Duration)
	Tick(d time.Duration) <-chan time.Time
	Ticker(d time.Duration) *Ticker
	Timer(d time.Duration) *Timer
	WithDeadline(parent context.Context, d time.Time) (context.Context, context.CancelFunc)
	WithTimeout(parent context.Context, t time.Duration) (context.Context, context.CancelFunc)
}

// New returns an instance of a real-time clock.
func New() Clock {
	return &realClock{}
}

type realClock struct{}

func (c *realClock) After(d time.Duration) <-chan time.Time { return time.After(d) }

func (c *realClock) AfterFunc(d time.Duration, f func()) *Timer {
	return &Timer{timer: time.AfterFunc(d, f)}
}

func (c *realClock) Now() time.Time                        { return time.Now() }
func (c *realClock) Since(t time.Time) time.Duration       { return time.Since(t) }
func (c *realClock) Until(t time.Time) time.Duration       { return time.Until(t) }
func (c *realClock) Sleep(d time.Duration)                 { time.Sleep(d) }
func (c *realClock) Tick(d time.Duration) <-chan time.Time { return time.Tick(d) }

func (c *realClock) Ticker(d time.Duration) *Ticker {
	t := time.NewTicker(d)
	return &Ticker{C: t.C, ticker: t}
}

func (c *realClock) Timer(d time.Duration) *Timer {
	t := time.NewTimer(d)
	return &Timer{C: t.C, timer: t}
}

func (c *realClock) WithDeadline(parent context.Context, d time.Time) (context.Context, context.CancelFunc) {
	return context.WithDeadline(parent, d)
}

func (c *realClock) WithTimeout(parent context.Context, t time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(parent, t)
}

// Timer represents a single event.
type Timer struct {
	C     <-chan time.Time
	timer *time.Timer
}

func (t *Timer) Stop() bool {
	if t.timer != nil {
		return t.timer.Stop()
	}
	return false
}

func (t *Timer) Reset(d time.Duration) bool {
	if t.timer != nil {
		return t.timer.Reset(d)
	}
	return false
}

// Ticker holds a channel that receives "ticks" at regular intervals.
type Ticker struct {
	C      <-chan time.Time
	ticker *time.Ticker
}

func (t *Ticker) Stop() {
	if t.ticker != nil {
		t.ticker.Stop()
	}
}

func (t *Ticker) Reset(dur time.Duration) {
	if t.ticker != nil {
		t.ticker.Reset(dur)
	}
}
