package home

import (
	"go-fiber-job/internal/vacancy"
	"go-fiber-job/pkg/tmpadapter"
	"go-fiber-job/views"
	"math"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *vacancy.VacancyRepository
}

type User struct {
	Id   int
	Name string
}

// конструктор
func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *vacancy.VacancyRepository) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
	}
	h.router.Get("/", h.home)
	h.router.Get("/404", h.error)
	h.router.Get("/login", h.login)

}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	PAGE_IMEMS := 2
	page := c.QueryInt("page", 1)
	count := h.repository.CountAll()
	vacancies, err := h.repository.GetAll(PAGE_IMEMS, (page-1)*PAGE_IMEMS)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	component := views.Main(vacancies, int(math.Ceil(float64(count/PAGE_IMEMS))), page)
	return tmpadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) login(c *fiber.Ctx) error {
	component := views.Login()
	return tmpadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	h.customLogger.Info().
		Bool("isAdmin", true).
		Str("email", "a@a.ru").
		Int("id", 10).
		Msg("Инфо")
	return c.SendString("Error")
}
