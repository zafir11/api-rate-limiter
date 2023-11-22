package middleware

import (
	"math"
	"time"
)

type Tokenbucket struct {
	tokens     float64
	maxtokens  float64
	refillRate float64
	lastRefill time.Time
}

func NewTokenbucket(maxtokens float64, refillRate float64) *Tokenbucket {
	return &Tokenbucket{
		tokens:     maxtokens,
		maxtokens:  maxtokens,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}
func (tb *Tokenbucket) refill() {
	now := time.Now()
	duration := now.Sub(tb.lastRefill)
	tokenToAdd := tb.refillRate * duration.Seconds()
	tb.tokens = math.Min(tb.tokens+tokenToAdd, tb.maxtokens)
	tb.lastRefill = now

}
func (tb *Tokenbucket) Request(tokens float64) bool {
	tb.refill()
	if tokens <= tb.tokens {
		tb.tokens -= tokens
		return true
	}
	return false
}
