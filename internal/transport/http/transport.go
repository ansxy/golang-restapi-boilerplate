package http_transport

import (
	"net/http"

	"github.com/ansxy/golang-boilerplate-gin/internal/middleware"
	v1 "github.com/ansxy/golang-boilerplate-gin/internal/transport/http/v1"
	custome_gin "github.com/ansxy/golang-boilerplate-gin/pkg/gin"
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
