package home

import (
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
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	// data := struct {
	// 	Count   int
	// 	IsAdmin bool
	// 	CanUse  bool
	// }{Count: 10, IsAdmin: true, CanUse: true}

	users := []User{
		{Id: 1, Name: "Anton"},
		{Id: 2, Name: "Vasya"},
		{Id: 3, Name: "Petya"},
	}
	names := []string{"Anton", "Vasia", "Petya"}
	data := struct {
		Names []string
		Users []User
	}{Names: names, Users: users}
	return c.Render("page", data)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Str("email", "a@a.ru").
		Int("id", 10).
		Msg("Инфо")
	return c.SendString("Eroor")
}
