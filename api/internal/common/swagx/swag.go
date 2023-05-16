package swagx

import (
	"bytes"
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
	"html/template"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/zeromicro/go-zero/core/logx"
	"pure-go-zero-admin/api/internal/common/errorx"
	"pure-go-zero-admin/api/internal/common/result"
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

// docsJson returns swagger json
func (c *FileConfig) docsJson() http.HandlerFunc {
	swaggerFile, err := os.Open(c.output + "/swagger.json")
	if err != nil {
		return c.docsErr(err)
	}
	defer swaggerFile.Close()

	swaggerByte, err := io.ReadAll(swaggerFile)
	if err != nil {
		return c.docsErr(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if _, err = w.Write(swaggerByte); err != nil {
			result.Response(w, r, err)
		}
	}
}

// docs returns swagger html
func (c *FileConfig) docs() http.HandlerFunc {
	logx.Info("swagger docs \n")
	fmt.Printf("swagger docs \n")
	tmpl := template.Must(template.New("swaggerdoc").Parse(swaggerTemplateV2))
	buf := bytes.NewBuffer(nil)
	err := tmpl.Execute(buf, c.swaggerConfig)
	if err != nil {
		return c.docsErr(err)
	}

	uiHTML := buf.Bytes()
	logx.Infof("swagger env: %s \n", c.env)
	fmt.Printf("swagger env: %s", c.env)
	return func(w http.ResponseWriter, r *http.Request) {
		// permission
		if c.env == service.ProMode {
			result.Response(w, r, errorx.NewError(http.StatusForbidden))
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if _, err = w.Write(uiHTML); err != nil {
			result.Response(w, r, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// docs returns swagger error
func (c *FileConfig) docsErr(err error) http.HandlerFunc {
	logx.Error(err)

	return func(w http.ResponseWriter, r *http.Request) {
		result.Response(w, r, errorx.NewError(http.StatusNotFound))
	}
}
