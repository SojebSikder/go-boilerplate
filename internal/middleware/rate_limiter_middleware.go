// middleware/rate_limiter.go
package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type limiterEntry struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu        sync.RWMutex
	visitors  = make(map[string]*limiterEntry)
	rateLimit = rate.Limit(1)
	burst     = 1
)

func getVisitor(ip string) *rate.Limiter {
	mu.RLock()
	v, exists := visitors[ip]
	mu.RUnlock()

	if exists {
		mu.Lock()
		v.lastSeen = time.Now()
		mu.Unlock()
		return v.limiter
	}

	limiter := rate.NewLimiter(rateLimit, burst)
	mu.Lock()
	visitors[ip] = &limiterEntry{limiter: limiter, lastSeen: time.Now()}
	mu.Unlock()

	return limiter
}

func startCleanup() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

func init() {
	go startCleanup()
}

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.GetHeader("X-Forwarded-For")
		if ip == "" {
			ip = c.ClientIP()
		}

		limiter := getVisitor(ip)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": "Too many requests. Please slow down.",
			})
			return
		}

		c.Next()
	}
}
