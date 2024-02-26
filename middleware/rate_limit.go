package middleware

import (
	"fmt"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	gogin "github.com/cntech-io/cntechkit-gogin/v2"
	errormessage "github.com/cntech-io/cntechkit-gogin/v2/error-message"
	"github.com/cntech-io/cntechkit-gogin/v2/response"
	"github.com/cntech-io/cntechkit-gogin/v2/utility"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func keyFunc(c *gin.Context) string {
	clientIp, _ := utility.GetClientIPFromContext(c)
	return clientIp
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	clientIp, _ := utility.GetClientIPFromContext(c)
	gogin.LOGGER.Warn(fmt.Sprintf("IP: %s rate limit exceeded", clientIp))

	c.JSON(
		response.New().TooManyRequests(
			errormessage.New("ERR_TOO_MANY_REQUEST", "Too many requests, please try again later."),
		),
	)
}

func NewRateLimitMiddleware(redisClient *redis.Client, rate time.Duration, limit uint) gin.HandlerFunc {
	rateLimitStore := ratelimit.RedisStore(
		&ratelimit.RedisOptions{
			RedisClient: redisClient,
			Rate:        rate,
			Limit:       limit,
		},
	)

	return ratelimit.RateLimiter(rateLimitStore, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})
}
