package middlewares

import (
	cntechkitgo "github.com/cntech-io/cntechkit-go"
	cntechkitgogin "github.com/cntech-io/cntechkit-gogin"
	endusermessage "github.com/cntech-io/cntechkit-gogin/constants/end_user_message"
	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		username, password, ok := context.Request.BasicAuth()
		if !ok || username == "" || password == "" {
			cntechkitgo.NewLogger(&cntechkitgo.LoggerConfig{
				AppName: "cntechkit-gogin",
			}).Error(string(endusermessage.BasicAuthInvalidParams))
			cntechkitgogin.
				NewResponse(string(endusermessage.BasicAuthInvalidParams)).
				BadRequest(context, cntechkitgogin.InvalidParams)
			context.Abort()
		}

		context.Next()
	}
}
