package public

import (
	"github.com/ansxy/golang-boilerplate-gin/internal/response"
	"github.com/gin-gonic/gin"
)

type Testing struct {
	Lorem string `json:"lorem" required:"true" validate:"required"`
	Ipsum string `json:"ipsum"  required:"true" validate:"required"`
}

func (h *PublicHandler) Health(c *gin.Context) {
	var json Testing
	err := h.v.ValidateStruct(c.Request, &json)
	if err != nil {
		response.Error(c.Writer, err)
		return
	}

	response.Success(c.Writer, nil)
}
