package utils

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sojebsikder/go-boilerplate/internal/config"
	"github.com/sojebsikder/go-boilerplate/pkg/redis"
)

type RateLimiter struct {
	redis *redis.Redis
	cfg   *config.Config
}

func NewRateLimiter(redisClient *redis.Redis, cfg *config.Config) *RateLimiter {
	return &RateLimiter{
		redis: redisClient,
		cfg:   cfg,
	}
}

func (rl *RateLimiter) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := fmt.Sprintf("rate_limit:%s", ip)

		ctx := context.Background()

		count, err := rl.redis.Client.Incr(ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Rate limiter internal error"})
			c.Abort()
			return
		}

		if count == 1 {
			// If it's the first request, set the TTL
			rl.redis.Client.Expire(ctx, key, rl.cfg.RateLimit.RateLimitDuration)
		}

		if count > int64(rl.cfg.RateLimit.RateLimitMaxRequests) {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}

		c.Next()
	}
}
