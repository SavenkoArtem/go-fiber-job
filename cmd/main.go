package main

import (
	"go-fiber-job/config"
	"go-fiber-job/internal/home"
	"go-fiber-job/pkg/logger"
	"strings"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	engine := html.New("./html", ".html") // в указанной дериктории парсит все файлы с выбранным расширением
	engine.AddFuncMap(map[string]interface{}{
		"ToUpper": func(c string) string {
			return strings.ToUpper(c)
		},
	})

	app := fiber.New(fiber.Config{
		Views: engine, // это temple engin который работает с шаблонам
	})
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New()) // позволяет восстановиться после падения (не падать)

	home.NewHandler(app, customLogger) // вызов handler-а internal/home

	app.Listen(":3000")
}
