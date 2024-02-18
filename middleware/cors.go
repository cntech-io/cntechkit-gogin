package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(config *cors.Config) gin.HandlerFunc {
	if config == nil {
		return cors.New(cors.DefaultConfig())
	}
	return cors.New(*config)
}
