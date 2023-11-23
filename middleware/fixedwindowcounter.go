package middleware

import (
	"sync"
	"time"
)

type FixedWindow struct {
	windowSize        time.Duration
	requestsThreshold float64
	maxRequest        float64
	resetTime         time.Time
	mu                sync.Mutex
}

func NewFixedWindow(windowSize time.Duration, maxRequest float64) *FixedWindow {
	return &FixedWindow{
		windowSize:        windowSize,
		requestsThreshold: 0,
		maxRequest:        maxRequest,
		resetTime:         time.Now().Add(windowSize),
	}
}
func (fw *FixedWindow) Allow() bool {
	fw.mu.Lock()
	defer fw.mu.Unlock()
	now := time.Now()
	// reset the counter if the pass the window
	if now.After(fw.resetTime) {
		fw.requestsThreshold = 0
		fw.resetTime = now.Add(fw.windowSize)
	}
	// check for max request
	if fw.requestsThreshold <= fw.maxRequest {
		fw.requestsThreshold++
		return true
	}
	return false
}
