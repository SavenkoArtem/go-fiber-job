package main

import (
	"go-fiber-job/config"
	"go-fiber-job/internal/home"
	"go-fiber-job/internal/sitemap"
	"go-fiber-job/internal/vacancy"
	"go-fiber-job/pkg/database"
	"go-fiber-job/pkg/logger"
	"go-fiber-job/pkg/middleware"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	dbConfig := config.NewDatabaseConfig()
	customLogger := logger.NewLogger(logConfig)

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New()) // позволяет восстановиться после падения (не падать)
	app.Static("/public", "./public")
	app.Static("/robots.txt", "./public/robots.txt")
	dbpool := database.CreateDbPool(dbConfig, customLogger)
	defer dbpool.Close()

	// Save session
	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	store := session.New(session.Config{
		Storage: storage,
	})
	app.Use(middleware.AuthMiddleware(store))

	// Repository
	vacancyRepo := vacancy.NewVacancyRepository(dbpool, customLogger)

	// Handler

	home.NewHandler(app, customLogger, vacancyRepo, store) // вызов handler-а internal/home
	vacancy.NewHandler(app, customLogger, vacancyRepo)
	sitemap.NewHandler(app)

	app.Listen(":3000")
}
