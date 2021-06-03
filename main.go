package main

import (
	"log"

	"github.com/y-yagi/go-api-template/database"
	"github.com/y-yagi/go-api-template/routes"
)

func main() {
	err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	app := routes.New()
	log.Fatal(app.Listen(":3000"))
}
