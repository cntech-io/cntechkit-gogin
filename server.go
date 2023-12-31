package cntechkitgogin

import (
	"fmt"

	cntechkitgo "github.com/cntech-io/cntechkit-go"
	"github.com/cntech-io/cntechkit-gogin/responses"
	"github.com/gin-gonic/gin"
)

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
	if s.config.AppPort == "" {
		s.router.Run(":8080")
		cntechkitgo.NewLogger(&cntechkitgo.LoggerConfig{
			AppName: "cntechkit-gogin",
		}).Info("Starting server on port 8080")
	} else {
		s.router.Run(s.config.AppPort)
		cntechkitgo.NewLogger(&cntechkitgo.LoggerConfig{
			AppName: "cntechkit-gogin",
		}).Info("Starting server on port " + s.config.AppPort)
	}
}

func (s *Server) Router() *gin.Engine {
	return s.router
}
