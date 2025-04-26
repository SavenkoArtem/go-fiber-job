package main

import (
	"go-fiber-job/config"
	"go-fiber-job/internal/home"
	"go-fiber-job/internal/vacancy"
	"go-fiber-job/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New()) // позволяет восстановиться после падения (не падать)
	app.Static("/public", "./public")

	home.NewHandler(app, customLogger) // вызов handler-а internal/home
	vacancy.NewHandler(app, customLogger)

	app.Listen(":3000")
}
