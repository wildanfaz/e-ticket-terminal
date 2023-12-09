package users

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/e-ticket-terminal/internal/helpers"
	"github.com/wildanfaz/e-ticket-terminal/internal/models"
)

func (s *Service) TopUp(c echo.Context) error {
	var (
		response = helpers.NewResponse()
		payload  models.User
	)

	err := c.Bind(&payload)
	if err != nil {
		s.log.Errorln("Invalid Payload :", err)
		return c.JSON(http.StatusBadRequest,
			response.AsError().
				WithMessage("Invalid Payload"))
	}

	err = validation.ValidateStruct(&payload,
		validation.Field(&payload.Balance, validation.Required, validation.Min(1)),
	)
	if err != nil {
		s.log.Errorln("Invalid Payload :", err)
		return c.JSON(http.StatusBadRequest,
			response.AsError().
				WithMessage(err.Error()))
	}

	payload.Email = c.Get("email").(string)

	err = s.usersRepo.TopUp(c.Request().Context(), payload)
	if err != nil {
		s.log.Errorln("Failed to top up user :", err)
		return c.JSON(http.StatusInternalServerError,
			response.AsError().
				WithMessage("Internal Server Error"))
	}

	s.log.Println("Top Up Success")
	return c.JSON(http.StatusOK,
		response.WithMessage("Top Up Success"),
	)
}
