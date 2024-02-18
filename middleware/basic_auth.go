package middleware

import (
	gogin "github.com/cntech-io/cntechkit-gogin/v2"
	errormessage "github.com/cntech-io/cntechkit-gogin/v2/error-message"
	"github.com/cntech-io/cntechkit-gogin/v2/response"
	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok || username == "" || password == "" {
			gogin.LOGGER.Warn("invalid basic auth header")
			c.JSON(response.New().BadRequest(errormessage.ERR_INVALID_BASIC_AUTH_HEADER))
			c.Abort()
		}
		c.Next()
	}
}
