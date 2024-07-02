package app

import custome_gin "github.com/ansxy/niaga-catering-be/pkg/gin"

func Exec() (err error) {

	app := custome_gin.NewGinApp()

	return app.Run(":8080")
}
