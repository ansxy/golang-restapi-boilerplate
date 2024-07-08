package app

import (
	"net/http"

	http_transport "github.com/ansxy/niaga-catering-be/internal/transport/http"
	custome_http "github.com/ansxy/niaga-catering-be/pkg/http"
)

func Exec() (err error) {
	handler := http_transport.NewHttpTransport()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err = custome_http.NewHttpServer(srv)
	if err != nil {
		return err
	}

	return
}
