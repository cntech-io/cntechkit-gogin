package cntechkitgogin

import (
	"fmt"

	cntechkitgo "github.com/cntech-io/cntechkit-go"
	"github.com/cntech-io/cntechkit-gogin/responses"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	AppPort    string
	CorsConfig *cors.Config
}

type Server struct {
	config *ServerConfig
	router *gin.Engine
}

var env = cntechkitgo.NewServerEnv()

func NewServer(conf ServerConfig) *Server {
	server := &Server{}
	server.config = &conf
	if env.DebugModeFlag {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	server.router = gin.New()

	server.router.SetTrustedProxies(env.TrustedProxies)

	return server
}

func (s *Server) AddController(c *Controller) *Server {
	if c.path == "" {
		panic("controller path missing")
	}
	if c.version == "" {
		panic("controller version missing")
	}
	group := s.router.Group(fmt.Sprintf("%v%v", c.version, c.path))
	for _, resource := range c.apis {
		group.Handle(string(resource.method), resource.path, append(resource.middlewares, resource.handler)...)
	}
	return s
}

func (s *Server) AttachMiddleWare(middleware gin.HandlerFunc) *Server {
	s.router.Use(middleware)
	return s
}

func (s *Server) AttachHealth() *Server {
	group := s.router.Group("/health")
	group.GET("/", func(c *gin.Context) {
		c.JSON(responses.New().OK("Healthy"))
	})
	return s
}

func (s *Server) Run() {
	var PORT = ":8080"
	if s.config.AppPort != "" {
		PORT = s.config.AppPort
	}

	s.router.Run(PORT)

	cntechkitgo.
		NewLogger(
			&cntechkitgo.LoggerConfig{
				AppName: "cntechkit-gogin",
			},
		).
		Info(fmt.Sprintf("Starting server on port: %v", PORT))
}

func (s *Server) Router() *gin.Engine {
	return s.router
}
