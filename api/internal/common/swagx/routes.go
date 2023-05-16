package swagx

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, c *FileConfig) {
	server.AddRoutes([]rest.Route{
		{
			Method:  http.MethodGet,
			Path:    "/swagger-json",
			Handler: c.docsJson(),
		},
		{
			Method:  http.MethodGet,
			Path:    "/swagger",
			Handler: c.docs(),
		},
	})
}
