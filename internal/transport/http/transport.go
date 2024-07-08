package http_transport

import (
	"net/http"

	v1 "github.com/ansxy/niaga-catering-be/internal/transport/http/v1"
	custome_gin "github.com/ansxy/niaga-catering-be/pkg/gin"
)

func NewHttpTransport() http.Handler {
	app := custome_gin.NewGinApp()

	v1.NewV1Transport(app.Group("/api"))

	return app
}
