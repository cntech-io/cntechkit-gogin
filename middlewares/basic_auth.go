package middlewares

import (
    cntechkitgo "github.com/cntech-io/cntechkit-go"
    "github.com/cntech-io/cntechkit-gogin/errorcodes"
    "github.com/cntech-io/cntechkit-gogin/responses"
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
            Error(string(errorcodes.ERRC17006.Default))
            c.JSON(responses.New().BadRequest(errorcodes.ERRC17006))
            c.Abort()
        }
        c.Next()
    }
}
