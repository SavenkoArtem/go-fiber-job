package home

import (
	"go-fiber-job/pkg/tmpadapter"
	"go-fiber-job/views"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

type User struct {
	Id   int
	Name string
}

// конструктор
func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	h.router.Get("/", h.home)
	h.router.Get("/404", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main()
	return tmpadapter.Render(c, component)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Str("email", "a@a.ru").
		Int("id", 10).
		Msg("Инфо")
	return c.SendString("Eroor")
}
