package middleware

import (
	"github.com/cntech-io/cntechkit-go/v2/logger"
	errormessage "github.com/cntech-io/cntechkit-gogin/v2/error-message"
	"github.com/cntech-io/cntechkit-gogin/v2/response"
	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok || username == "" || password == "" {
			logger.
				NewLogger(
					&logger.LoggerConfig{
						AppName: "cntechkit-gogin",
					},
				).
				Error(string(errormessage.ERR_INVALID_BASIC_AUTH_HEADER_FORMAT.Code))
			c.JSON(response.New().BadRequest(errormessage.ERR_INVALID_BASIC_AUTH_HEADER_FORMAT))
			c.Abort()
		}
		c.Next()
	}
}
