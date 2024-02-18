package server

import (
	"fmt"

	e "github.com/cntech-io/cntechkit-go/v2/env"
	gogin "github.com/cntech-io/cntechkit-gogin/v2"
	"github.com/cntech-io/cntechkit-gogin/v2/controller"
	errormessage "github.com/cntech-io/cntechkit-gogin/v2/error-message"
	"github.com/cntech-io/cntechkit-gogin/v2/response"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

var env = e.NewServerEnv()

func NewServer() *Server {
	server := &Server{}

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
		panic("Controller path missing")
	}
	if c.GetVersion() == "" {
		panic("Controller version missing")
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

func (s *Server) AttachErrorDefinitions() *Server {
	group := s.router.Group("/error-definitions")
	group.GET("/", func(c *gin.Context) {
		c.JSON(response.New().WithData("success", errormessage.List))
	})
	return s
}

func (s *Server) Run() {
	var PORT = ":8080"
	if env.AppPort != "" {
		PORT = env.AppPort
	}

	err := s.router.Run(PORT)
	if err != nil {
		panic(fmt.Sprintf("Error starting server: %v", err))
	}

	gogin.LOGGER.Info(fmt.Sprintf("Starting server on port: %v", PORT))
}

func (s *Server) Router() *gin.Engine {
	return s.router
}
