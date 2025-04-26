package vacancy

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHander struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

type Vacancy struct {
	role        string
	company     string
	typeVacancy string
	salary      string
	location    string
	email       string
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &VacancyHander{
		router:       router,
		customLogger: customLogger,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHander) createVacancy(c *fiber.Ctx) error {
	return c.SendString("CreateVacancy")
}
