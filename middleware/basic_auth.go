package middleware

import (
	cntechkitgo "github.com/cntech-io/cntechkit-go"
	errormessage "github.com/cntech-io/cntechkit-gogin/error-message"
	"github.com/cntech-io/cntechkit-gogin/response"
	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok || username == "" || password == "" {
			cntechkitgo.
				NewLogger(
					&cntechkitgo.LoggerConfig{
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