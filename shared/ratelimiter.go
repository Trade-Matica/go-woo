package shared

import (
	"sync"
	"time"
)

type RateLimiter struct {
	limit         int
	duration      time.Duration
	requests      int
	lastResetTime time.Time
	mutex         sync.Mutex
}

func NewRateLimiter(limit int, duration time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:         limit,
		duration:      duration,
		requests:      0,
		lastResetTime: time.Now(),
	}
}

func (r *RateLimiter) Wait() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if time.Since(r.lastResetTime) >= r.duration {
		r.requests = 0
		r.lastResetTime = time.Now()
	}

	for r.requests >= r.limit {
		time.Sleep(time.Millisecond * 1)
	}

	r.requests++
}
