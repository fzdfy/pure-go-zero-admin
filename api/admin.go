package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"pure-go-zero-admin/api/internal/common/errorx"
	"pure-go-zero-admin/api/internal/common/swagx"

	"pure-go-zero-admin/api/internal/config"
	"pure-go-zero-admin/api/internal/handler"
	"pure-go-zero-admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "api/etc/admin-api.yaml", "the config file")

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host
//	@BasePath	/_

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("c.Mode: %s \n", c.Mode)

	swagConfig := swagx.NewConfig(swagx.WithGeneralInfo("api/admin.go"), swagx.WithOutput("api/docs"), swagx.WithEnv(c.Mode))
	swagConfig.Makefile()
	swagx.RegisterHandlers(server, swagConfig)
	fmt.Printf("swagx.RegisterHandlers \n")

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
