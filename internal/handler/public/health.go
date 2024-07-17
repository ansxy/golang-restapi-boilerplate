package public

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *PublicHandler) Health(c *gin.Context) {
	time.Sleep(10 * time.Second) // Adjust the delay as needed
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})

}
