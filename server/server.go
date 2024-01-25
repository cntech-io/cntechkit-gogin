package server

import (
	"fmt"

	e "github.com/cntech-io/cntechkit-go/v2/env"
	"github.com/cntech-io/cntechkit-go/v2/logger"
	"github.com/cntech-io/cntechkit-gogin/v2/controller"
	"github.com/cntech-io/cntechkit-gogin/v2/response"
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

var env = e.NewServerEnv()

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

func (s *Server) AddController(c *controller.Controller) *Server {
	if c.GetPath() == "" {
		panic("controller path missing")
	}
	if c.GetVersion() == "" {
		panic("controller version missing")
	}
	group := s.router.Group(fmt.Sprintf("%v%v", c.GetVersion(), c.GetPath()))
	for _, api := range c.GetApis() {
		group.Handle(
			string(api.GetMethod()),
			api.GetPath(),
			append(api.GetMiddlewares(), api.GetHandler())...,
		)
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
		c.JSON(response.New().OK("Healthy"))
	})
	return s
}

func (s *Server) Run() {
	var PORT = ":8080"
	if s.config.AppPort != "" {
		PORT = s.config.AppPort
	}

	s.router.Run(PORT)

	logger.
		NewLogger(
			&logger.LoggerConfig{
				AppName: "cntechkit-gogin",
			},
		).
		Info(fmt.Sprintf("Starting server on port: %v", PORT))
}

func (s *Server) Router() *gin.Engine {
	return s.router
}
