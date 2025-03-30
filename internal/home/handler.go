package home

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {
	router fiber.Router
}

// конструктор
func NewHandler(router fiber.Router) {
	h := &HomeHandler{
		router: router,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Eroor")
}
