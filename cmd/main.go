package main

import (
	"log"

	"github.com/ansxy/niaga-catering-be/cmd/app"
)

func main() {
	err := app.Exec()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
