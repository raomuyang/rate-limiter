package rate_limiter

import (
	"errors"
	"time"
)

type RateLimiter interface {
	SetRate(permitsPerSecond int)

	GetRate() int

	Acquire(permits int) int64

	TryAcquire(permits int, timeout time.Duration) bool

}

type Stopwatch struct {
	startNanos   int64
	elapsedNanos int64
	isRun        bool
}

func (stopwatch *Stopwatch) Start() error {
	if stopwatch.isRun {
		return errors.New("this stopwatch is already running")
	}
	stopwatch.isRun = true
	stopwatch.startNanos = time.Now().UnixNano()
	return nil
}

func (stopwatch *Stopwatch) Stop() error {
	if !stopwatch.isRun {
		return errors.New("this stopwatch is already stopped")
	}
	stopwatch.isRun = false
	stopwatch.elapsedNanos += time.Now().UnixNano() - stopwatch.startNanos
	return nil
}

func (stopwatch *Stopwatch) ElapsedNanos() int64 {
	if stopwatch.isRun {
		return time.Now().UnixNano() - stopwatch.startNanos
	}
	return stopwatch.elapsedNanos
}

func (stopwatch *Stopwatch) Reset() {
	stopwatch.elapsedNanos = 0
	stopwatch.isRun = false
}

func CreateStartedStopwatch() *Stopwatch {
	stopwatch := Stopwatch{}
	stopwatch.Start()
	return &stopwatch
}