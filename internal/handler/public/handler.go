package public

import (
	custom_validator "github.com/ansxy/golang-boilerplate-gin/pkg/validator"
	"github.com/gin-gonic/gin"
)

type PublicHandler struct {
	v custom_validator.Validator
}

func NewPublicHandler(route *gin.RouterGroup, validator custom_validator.Validator) *gin.RouterGroup {

	routeGroup := route.Group("/public")
	handler := &PublicHandler{
		v: validator,
	}
	routeGroup.POST("/health", handler.Health)

	return route
}
