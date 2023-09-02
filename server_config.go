package cntechkitgogin

import "github.com/gin-contrib/cors"

type ServerConfig struct {
	AppPort    string
	CorsConfig *cors.Config
}
