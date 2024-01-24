package api

import "github.com/gin-gonic/gin"

type RouteMethod string

const (
	GET    RouteMethod = "GET"
	POST   RouteMethod = "POST"
	PUT    RouteMethod = "PUT"
	PATCH  RouteMethod = "PATCH"
	DELETE RouteMethod = "DELETE"
)

type Api struct {
	path        string
	method      RouteMethod
	handler     gin.HandlerFunc
	middlewares []gin.HandlerFunc
}

func NewAPI(method RouteMethod, path string) *Api {
	return &Api{
		path:   path,
		method: method,
	}
}

func (a *Api) Handler(h gin.HandlerFunc) *Api {
	a.handler = h
	return a
}

func (a *Api) AddMiddleware(m gin.HandlerFunc) *Api {
	a.middlewares = append(a.middlewares, m)
	return a
}

func (a *Api) GetPath() string {
	return a.path
}

func (a *Api) GetMethod() RouteMethod {
	return a.method
}

func (a *Api) GetHandler() gin.HandlerFunc {
	return a.handler
}

func (a *Api) GetMiddlewares() []gin.HandlerFunc {
	return a.middlewares
}
