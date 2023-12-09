package terminals

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/e-ticket-terminal/internal/constants"
	"github.com/wildanfaz/e-ticket-terminal/internal/helpers"
	"github.com/wildanfaz/e-ticket-terminal/internal/models"
)

func (s *Service) UpdateTransaction(c echo.Context) error {
	var (
		response = helpers.NewResponse()
		payload  models.UpdateTransaction
	)

	err := c.Bind(&payload)
	if err != nil {
		s.log.Errorln("Invalid Payload :", err)
		return c.JSON(http.StatusBadRequest,
			response.AsError().
				WithMessage("Invalid Payload"))
	}

	err = validation.ValidateStruct(&payload,
		validation.Field(&payload.ToTerminalId, validation.Required, validation.Min(1)),
	)
	if err != nil {
		s.log.Errorln("Invalid Payload :", err)
		return c.JSON(http.StatusBadRequest,
			response.AsError().
				WithMessage(err.Error()))
	}

	email := c.Get("email").(string)
	user, err := s.usersRepo.GetUserByEmail(c.Request().Context(), email)
	if err != nil {
		s.log.Errorln("Failed to get user :", err)
		return c.JSON(http.StatusInternalServerError,
			response.AsError().
				WithMessage("Internal Server Error"))
	}

	if user == nil {
		s.log.Errorln("User not found")
		return c.JSON(http.StatusInternalServerError,
			response.AsError().
				WithMessage("Internal Server Error"))
	}

	err = s.terminalsRepo.UpdateTransaction(c.Request().Context(), payload, *user)
	if err != nil && err.Error() != constants.InsufficientBalance {
		s.log.Errorln("Failed to update transaction :", err)
		return c.JSON(http.StatusInternalServerError,
			response.AsError().
				WithMessage("Internal Server Error"))
	}

	if err != nil && err.Error() == constants.InsufficientBalance {
		s.log.Errorln("Failed to update transaction :", err)
		return c.JSON(http.StatusBadRequest,
			response.AsError().
				WithMessage(err.Error()))
	}

	s.log.Println("Update Transaction Success")
	return c.JSON(http.StatusOK,
		response.WithMessage("Update Transaction Success"),
	)
}
