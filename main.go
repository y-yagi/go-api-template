package main

import (
	"log"

	"github.com/y-yagi/go-api-template/database"
	"github.com/y-yagi/go-api-template/routes"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	logger := log.Default()
	logger.SetOutput(&lumberjack.Logger{
		Filename:   "log/application.log",
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     3, //days
		Compress:   true,
	})

	err := database.New(logger)
	if err != nil {
		log.Fatal(err)
	}

	app := routes.New()
	log.Fatal(app.Listen(":3000"))
}
