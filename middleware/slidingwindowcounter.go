package middleware

import (
	"sync"
	"time"
)

type slidingWindow struct {
	windowSize          time.Duration
	requestRate         float64
	currentWindowCount  float64
	previousWindowCount float64
	resetTime           time.Time
	mu                  sync.Mutex
}

func NewSlidingWindow(windowSize time.Duration, requestRate float64) *slidingWindow {

	return &slidingWindow{
		windowSize:          windowSize,
		requestRate:         requestRate,
		currentWindowCount:  0,
		previousWindowCount: 0,
		resetTime:           time.Now().Add(windowSize),
	}
}
func (sw *slidingWindow) Allow() bool {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	now := time.Now()
	if now.After(sw.resetTime) {
		sw.previousWindowCount = sw.currentWindowCount
		sw.currentWindowCount = 0
		sw.resetTime = now.Add(sw.windowSize)
	}
	slidingWindowCount := (float64(sw.windowSize - (sw.resetTime.Sub(now)))) / float64(sw.windowSize) * sw.previousWindowCount
	slidingWindowCount += (float64(sw.resetTime.Sub(now))) / float64(sw.windowSize) * sw.currentWindowCount
	if slidingWindowCount <= sw.requestRate {
		sw.currentWindowCount++
		return true
	}
	return false
}
