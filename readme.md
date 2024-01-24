## Methods

| Method                                              | Description                                          |
| --------------------------------------------------- | ---------------------------------------------------- |
| NewServer(conf ServerConfig)                        | Creates server (gin framework)                       |
| server.AddController(c \*Controller)                | Adds controller to server                            |
| server.AttachMiddleWare(middleware gin.HandlerFunc) | Attachs middleware                                   |
| server.AttachHealth()                               | Attachs /health api                                  |
| server.Run()                                        | Runs server                                          |
| server.Router()                                     | Gets router ( \*gin.Engine)                          |
| NewController(version string, path string)          | Creates controller                                   |
| controller.AttachAPI(a \*api)                       | Attachs Api to controller                            |
| NewAPI(method RouteMethod, path string)             | Creates Api                                          |
| api.Handler(h gin.HandlerFunc)                      | Adds handler to Api                                  |
| api.AddMiddleware(m gin.HandlerFunc)                | Adds middleware to Api                               |
| NewAPI(method RouteMethod, path string)             | Creates Api                                          |
| responses.New()                                     | Creates Response                                     |
| response.BadRequest(errcode errorcodes.ErrorCode)   | Adds BadRequest to Response with predefined errors   |
| response.Unauthorized(errcode errorcodes.ErrorCode) | Adds Unauthorized to Response with predefined errors |
| response.NotFound(errcode errorcodes.ErrorCode)     | Adds NotFound to Response with predefined errors     |
| response.NoContent(message string)                  | Adds NoContent to Response with message              |
| response.OK(message string)                         | Adds OK Response to with message                     |
