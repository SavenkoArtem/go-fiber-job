package vacancy

import (
	"go-fiber-job/pkg/tmpadapter"
	"go-fiber-job/pkg/validator"
	"go-fiber-job/views/components"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHander struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *VacancyRepository
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *VacancyRepository) {
	h := &VacancyHander{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (h *VacancyHander) createVacancy(c *fiber.Ctx) error {
	form := VacancyCreateForm{
		Role:     c.FormValue("role"),
		Company:  c.FormValue("company"),
		Type:     c.FormValue("type"),
		Salary:   c.FormValue("salary"),
		Location: c.FormValue("location"),
		Email:    c.FormValue("email"),
	}
	errors := validate.Validate(
		&validators.StringIsPresent{Name: "Role", Field: form.Role, Message: "Должность не указана"},
		&validators.StringIsPresent{Name: "Company", Field: form.Company, Message: "Компания не указана"},
		&validators.StringIsPresent{Name: "Type", Field: form.Type, Message: "Тип не указан"},
		&validators.StringIsPresent{Name: "Salary", Field: form.Salary, Message: "Зарплата не указана"},
		&validators.StringIsPresent{Name: "Location", Field: form.Location, Message: "Расположение не указано"},
		&validators.EmailIsPresent{Name: "Email", Field: form.Email, Message: "Email не задан или неверный"},
	)
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tmpadapter.Render(c, component, http.StatusBadRequest)
	}

	err := h.repository.addVacancy(form)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		component = components.Notification("Error, please try later", components.NotificationFail)
		return tmpadapter.Render(c, component, http.StatusBadRequest)
	}
	component = components.Notification("The vacancy was successfully created", components.NotificationSuccess)
	return tmpadapter.Render(c, component, http.StatusOK)
}
