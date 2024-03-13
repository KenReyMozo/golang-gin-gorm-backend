package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/juju/ratelimit"

	"github.com/gin-gonic/gin"
)

func RateLimiter(capacity int64, fillInterval time.Duration) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, capacity, capacity)

	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too Many Requests",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

type RateLimiterConfig struct {
	requestCount  int64
	interval      time.Duration
	tokens        map[string]int64
	lastAccess    map[string]time.Time
	mu            sync.Mutex
}


func (limiter *RateLimiterConfig) AcceptRequest(clientIP string) bool {
	limiter.mu.Lock()
	defer limiter.mu.Unlock()

	// Check if client is in tokens map
	token, ok := limiter.tokens[clientIP]
	if !ok {
		// If client is not in tokens map, add them with a full token
		limiter.tokens[clientIP] = limiter.requestCount
		limiter.lastAccess[clientIP] = time.Now()
		return true
	}

	// Check if 3 seconds have passed since last access
	lastAccessTime := limiter.lastAccess[clientIP]
	if time.Since(lastAccessTime) >= limiter.interval {
		// If 3 seconds have passed, reset token and update last access time
		limiter.tokens[clientIP] = limiter.requestCount
		limiter.lastAccess[clientIP] = time.Now()
		return true
	}

	// Check if token is available
	if token > 0 {
		// Decrement token and update last access time
		limiter.tokens[clientIP]--
		limiter.lastAccess[clientIP] = time.Now()
		return true
	}

	return false
}

func NewRateLimiter(requestCount int64, interval time.Duration) *RateLimiterConfig {
	return &RateLimiterConfig{
		requestCount: requestCount,
		interval:     interval,
		tokens:       make(map[string]int64),
		lastAccess:   make(map[string]time.Time),
	}
}

func CustomRateLimiter(requestCount int64, interval time.Duration) gin.HandlerFunc {
	limiter := NewRateLimiter(requestCount, interval)

	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		if limiter.AcceptRequest(clientIP) {
			c.Next()
		} else {
			c.String(http.StatusTooManyRequests, "Too many requests")
			c.Abort()
		}
	}
}