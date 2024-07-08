package v1

import (
	"github.com/ansxy/golang-boilerplate-gin/internal/handler/public"
	"github.com/gin-gonic/gin"
)

func NewV1Transport(route *gin.RouterGroup) *gin.RouterGroup {
	routeGroup := route.Group("/v1")

	public.NewPublicHandler(routeGroup)

	return route
}
