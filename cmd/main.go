package main

import (
	"log"

	"github.com/ansxy/golang-boilerplate-gin/cmd/app"
)

func main() {
	err := app.Exec()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
