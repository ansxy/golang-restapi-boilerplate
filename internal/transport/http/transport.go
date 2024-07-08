package http_transport

import (
	"net/http"

	"github.com/ansxy/niaga-catering-be/internal/middleware"
	v1 "github.com/ansxy/niaga-catering-be/internal/transport/http/v1"
	custome_gin "github.com/ansxy/niaga-catering-be/pkg/gin"
	"github.com/sirupsen/logrus"
)

func NewHttpTransport() http.Handler {
	app := custome_gin.NewGinApp()
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	app.Use(middleware.NewLogger(logger))

	v1.NewV1Transport(app.Group("/api"))

	return app
}
