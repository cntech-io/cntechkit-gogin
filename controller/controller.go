package controller

import "github.com/cntech-io/cntechkit-gogin/v2/api"

type Controller struct {
	version string
	path    string
	apis    []api.Api
}

func NewController(version string, path string) *Controller {
	return &Controller{
		version: version,
		path:    path,
	}
}

func (c *Controller) AttachAPI(a *api.Api) *Controller {
	c.apis = append(c.apis, *a)
	return c
}

func (c *Controller) GetPath() string {
	return c.path
}

func (c *Controller) GetVersion() string {
	return c.version
}

func (c *Controller) GetApis() []api.Api {
	return c.apis
}
