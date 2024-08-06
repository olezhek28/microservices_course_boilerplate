package user_v1

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"strings"
)

//go:embed swagger
var swaggerFS embed.FS

// SwaggerFS Обертка над файловой системой для http-сервера
// использует встроенную директорию swagger через embed
// при обращении к файлу-схеме (api.swagger.json) подменяет hostname на валидный
type SwaggerFS struct {
	handler  http.Handler
	httpPort string
}

// NewSwaggerFS новый экземпляр
func NewSwaggerFS(httpPort string) *SwaggerFS {
	root, err := fs.Sub(swaggerFS, "swagger")
	if err != nil {
		panic(err)
	}

	httpFs := http.FS(root)
	return &SwaggerFS{
		handler:  http.FileServer(httpFs),
		httpPort: httpPort,
	}
}

// ServeHTTP обертка над стандартным методом, обрабатывает обращение к api.swagger.json
func (f *SwaggerFS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api.swagger.json" {
		swaggerFile, err := swaggerFS.ReadFile("swagger/api.swagger.json")
		if err != nil {
			fmt.Println(err)
			return
		}

		hostName := strings.Split(r.Host, ":")[0]

		bt := bytes.Replace(swaggerFile, []byte("HOST_PLACEHOLDER"), []byte(net.JoinHostPort(hostName, f.httpPort)), 1)
		_, _ = w.Write(bt)
		return
	}

	f.handler.ServeHTTP(w, r)
}
