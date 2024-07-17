package app

import (
	"github.com/ansxy/golang-boilerplate-gin/config"
	http_transport "github.com/ansxy/golang-boilerplate-gin/internal/transport/http"
	custom_http "github.com/ansxy/golang-boilerplate-gin/pkg/http"
)

func Exec() (err error) {
	handler := http_transport.NewHttpTransport()
	cnf := config.NewConfig()

	err = custom_http.NewHttpServer(cnf, handler)
	if err != nil {
		return err
	}

	return
}
