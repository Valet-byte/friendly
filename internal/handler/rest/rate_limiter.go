package restHandler

import (
	"friendly/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReteLimiter struct {
	MaxRequestPerMinutes int
	RedisTraficService   *service.RedisService
}

func NewRateLimiter(maxRequestPerMinutes int, redisTraficService *service.RedisService) *ReteLimiter {
	return &ReteLimiter{MaxRequestPerMinutes: maxRequestPerMinutes, RedisTraficService: redisTraficService}
}

func (rl *ReteLimiter) RateLimiterMiddleware(c *gin.Context) {
	ip := c.ClientIP()
	err := rl.RedisTraficService.IncrementRequestCount(ip)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	requestCount, err := rl.RedisTraficService.GetRequestCount(ip)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	if requestCount > rl.MaxRequestPerMinutes {
		NewErrorResponse(c, http.StatusTooManyRequests, "Limit the number of requests per minute")
		c.Abort()
		return
	}

	c.Next()
}
