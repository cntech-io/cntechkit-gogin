### Server

| Method                                                                | Description                    |
| --------------------------------------------------------------------- | ------------------------------ |
| server.NewServer(conf ServerConfig)                                   | Creates server (gin framework) |
| &nbsp;&nbsp;&nbsp;&nbsp;.AddController(c \*Controller)                | Adds controller to server      |
| &nbsp;&nbsp;&nbsp;&nbsp;.AttachMiddleWare(middleware gin.HandlerFunc) | Attachs middleware             |
| &nbsp;&nbsp;&nbsp;&nbsp;.AttachHealth()                               | Attachs /health api            |
| &nbsp;&nbsp;&nbsp;&nbsp;.Run()                                        | Runs server                    |
| &nbsp;&nbsp;&nbsp;&nbsp;.Router()                                     | returns router ( \*gin.Engine) |

### Controller
| Method                                                              | Description                                          |
| ------------------------------------------------------------------- | ---------------------------------------------------- |
| controller.NewController(version string, path string)               | Creates controller                                   |
| &nbsp;&nbsp;&nbsp;&nbsp;.AttachAPI(a \*api)                         | Attachs Api to controller                            |

### Api
| Method                                                              | Description                                          |
| ------------------------------------------------------------------- | ---------------------------------------------------- |
| api.NewAPI(method RouteMethod, path string)                         | Creates Api                                          |
| &nbsp;&nbsp;&nbsp;&nbsp;.Handler(h gin.HandlerFunc)                 | Adds handler to Api                                  |
| &nbsp;&nbsp;&nbsp;&nbsp;.AddMiddleware(m gin.HandlerFunc)           | Adds middleware to Api                               |

### Response
| Method                                                              | Description                                          |
| ------------------------------------------------------------------- | ---------------------------------------------------- |
| response.New()                                                      | Creates Response                                     |
| &nbsp;&nbsp;&nbsp;&nbsp;.BadRequest(errcode errorcodes.ErrorCode)   | Adds BadRequest to Response with predefined errors   |
| &nbsp;&nbsp;&nbsp;&nbsp;.Unauthorized(errcode errorcodes.ErrorCode) | Adds Unauthorized to Response with predefined errors |
| &nbsp;&nbsp;&nbsp;&nbsp;.NotFound(errcode errorcodes.ErrorCode)     | Adds NotFound to Response with predefined errors     |
| &nbsp;&nbsp;&nbsp;&nbsp;.NoContent(message string)                  | Adds NoContent to Response with message              |
| &nbsp;&nbsp;&nbsp;&nbsp;.OK(message string)                         | Adds OK Response to with message                     |
