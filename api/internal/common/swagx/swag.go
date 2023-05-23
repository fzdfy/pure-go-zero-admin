package swagx

import (
	"os/exec"

	"github.com/zeromicro/go-zero/core/logx"
)

// FileConfig config
type FileConfig struct {
	// Go file path in which 'swagger general API Info' is written (default: "main.go").
	generalInfo string
	// Output types of generated files (docs.go, swagger.json, swagger.yaml) like go,json,yaml (default: "json").
	outputTypes string
	// Output directory for all the generated files(swagger.json, swagger.yaml and docs.go) (default: "./docs").
	output string
	// System env, if it's pro, do not display HTML (default: "dev").
	env string
	// swagger config, generally no setting is required.
	swaggerConfig
}

// swaggerConfig config
type swaggerConfig struct {
	// SpecURL the url to find the spec for, defaults to : swagger-json.
	SpecURL string
	// SwaggerHost for the js that generates the swagger ui site, defaults to: http://petstore.swagger.io/.
	SwaggerHost string
}

// Makefile is to create swagger file func
func (c *FileConfig) Makefile() {
	_ = exec.Command("swag", "fmt")
	cmd := exec.Command("swag", "init", "-requiredByDefault", "-g", c.generalInfo, "-ot", c.outputTypes, "-o", c.output)

	if err := cmd.Run(); err != nil {
		logx.Errorf("create swagger.json err : %s", err)
	}
}
