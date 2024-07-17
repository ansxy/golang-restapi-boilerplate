package custom_gin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewGinApp() *gin.Engine {
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	}))

	//implemenet Logger

	return app

}
