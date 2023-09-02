package cntechkitgogin

type Controller struct {
	version string
	path    string
	apis    []api
}

func NewController(version string, path string) *Controller {
	return &Controller{
		version: version,
		path:    path,
	}
}

func (c *Controller) AttachResource(a *api) *Controller {
	c.apis = append(c.apis, *a)
	return c
}
