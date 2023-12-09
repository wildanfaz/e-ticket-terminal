package terminals

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/e-ticket-terminal/internal/helpers"
	"github.com/wildanfaz/e-ticket-terminal/internal/models"
)

func (s *Service) AddTerminal(c echo.Context) error {
	var (
		response = helpers.NewResponse()
		payload  models.Terminal
	)

	err := c.Bind(&payload)
	if err != nil {
		s.log.Errorln("Invalid Payload :", err)
		return c.JSON(http.StatusBadRequest,
			response.AsError().
				WithMessage("Invalid Payload"))
	}

	err = validation.ValidateStruct(&payload,
		validation.Field(&payload.Name, validation.Required),
		validation.Field(&payload.LocationId, validation.Required, validation.Min(1)),
	)
	if err != nil {
		s.log.Errorln("Invalid Payload :", err)
		return c.JSON(http.StatusBadRequest,
			response.AsError().
				WithMessage(err.Error()))
	}

	err = s.terminalsRepo.AddTerminal(c.Request().Context(), payload)
	if err != nil {
		s.log.Errorln("Failed to add terminal :", err)
		return c.JSON(http.StatusInternalServerError,
			response.AsError().
				WithMessage("Internal Server Error"))
	}

	s.log.Println("Add Terminal Success")
	return c.JSON(http.StatusOK,
		response.WithMessage("Add Terminal Success"),
	)
}
