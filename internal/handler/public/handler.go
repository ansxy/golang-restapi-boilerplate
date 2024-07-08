package public

import (
	"github.com/gin-gonic/gin"
)

type PublicHandler struct {
}

func NewPublicHandler(route *gin.RouterGroup) *gin.RouterGroup {
	routeGroup := route.Group("/public")

	handler := &PublicHandler{}

	routeGroup.GET("/health", handler.Health)

	return route
}
