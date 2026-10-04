package utils

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	goredis "github.com/go-redis/redis/v8"
	"github.com/sojebsikder/go-boilerplate/internal/config"
	"github.com/sojebsikder/go-boilerplate/pkg/redis"
	"go.uber.org/zap"
)

var incrScript = goredis.NewScript(`
local count = redis.call("INCR", KEYS[1])
if count == 1 then
	redis.call("PEXPIRE", KEYS[1], ARGV[1])
end
local ttl = redis.call("PTTL", KEYS[1])
if ttl < 0 then
	redis.call("PEXPIRE", KEYS[1], ARGV[1])
	ttl = tonumber(ARGV[1])
end
return {count, ttl}
`)

const redisTimeout = 500 * time.Millisecond

type RateLimiter struct {
	redis  *redis.Redis
	cfg    *config.Config
	logger *zap.Logger
}

func NewRateLimiter(redisClient *redis.Redis, cfg *config.Config, logger *zap.Logger) *RateLimiter {
	return &RateLimiter{
		redis:  redisClient,
		cfg:    cfg,
		logger: logger,
	}
}

func (rl *RateLimiter) Limit() gin.HandlerFunc {
	limit := rl.cfg.RateLimit.RateLimitMaxRequests
	window := rl.cfg.RateLimit.RateLimitDuration

	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := fmt.Sprintf("rate_limit:%s", ip)

		ctx, cancel := context.WithTimeout(c.Request.Context(), redisTimeout)
		defer cancel()

		res, err := incrScript.Run(ctx, rl.redis.Client, []string{key}, window.Milliseconds()).Int64Slice()
		if err != nil {
			rl.logger.Error("rate_limiter_redis_error", zap.String("ip", ip), zap.Error(err))
			c.Next()
			return
		}

		count, ttl := res[0], time.Duration(res[1])*time.Millisecond

		c.Header("X-RateLimit-Limit", strconv.FormatInt(limit, 10))
		c.Header("X-RateLimit-Remaining", strconv.FormatInt(max(limit-count, 0), 10))

		if count > limit {
			c.Header("Retry-After", strconv.Itoa(int(math.Ceil(ttl.Seconds()))))
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			return
		}

		c.Next()
	}
}
