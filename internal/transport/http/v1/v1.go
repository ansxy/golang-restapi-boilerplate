package v1

import (
	"github.com/ansxy/niaga-catering-be/internal/handler/public"
	"github.com/gin-gonic/gin"
)

func NewV1Transport(route *gin.RouterGroup) *gin.RouterGroup {
	routeGroup := route.Group("/v1")

	public.NewPublicHandler(routeGroup)

	return route
}
