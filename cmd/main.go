package main

import (
	"go-fiber-job/config"
	"go-fiber-job/internal/home"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	app := fiber.New()
	app.Use(recover.New()) // позволяет восстановиться после падения (не падать)

	home.NewHandler(app) // вызов handler-а internal/home

	app.Listen(":3000")
}
