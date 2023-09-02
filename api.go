package cntechkitgogin

import "github.com/gin-gonic/gin"

type RouteMethod string

const (
	GET    RouteMethod = "GET"
	POST   RouteMethod = "POST"
	PUT    RouteMethod = "PUT"
	PATCH  RouteMethod = "PATCH"
	DELETE RouteMethod = "DELETE"
)

type api struct {
	path        string
	method      RouteMethod
	handler     gin.HandlerFunc
	middlewares []gin.HandlerFunc
}

func NewAPI(method RouteMethod, path string) *api {
	return &api{
		path:   path,
		method: method,
	}
}

func (a *api) Handler(h gin.HandlerFunc) *api {
	a.handler = h
	return a
}

func (a *api) AddMiddleware(m gin.HandlerFunc) *api {
	a.middlewares = append(a.middlewares, m)
	return a
}
