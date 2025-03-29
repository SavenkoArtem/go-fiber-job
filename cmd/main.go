package main

import (
	"go-fiber-job/internal/home"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	home.NewHandler(app) // вызов handler-а internal/home

	app.Listen(":3000")
}
