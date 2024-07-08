package public

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *PublicHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pang",
	})
}
