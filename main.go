package main

import (
	"log"

	"github.com/y-yagi/go-api-template/routes"
)

func main() {
	app := routes.New()
	log.Fatal(app.Listen(":3000"))
}
