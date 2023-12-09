package users

import (
	"net/http"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/e-ticket-terminal/internal/helpers"
	"github.com/wildanfaz/e-ticket-terminal/internal/models"
	"github.com/wildanfaz/e-ticket-terminal/internal/pkg"
)

func (s *Service) Register(c echo.Context) error {
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
		validation.Field(&payload.Email, validation.Required, is.Email),
		validation.Field(&payload.Password,
			validation.Required,
			validation.Length(6, 100),
			validation.Match(regexp.MustCompile("^[a-zA-Z0-9!@#$%^&*()-_+=]+$")),
		),
	)
	if err != nil {
		s.log.Errorln("Invalid Payload :", err)
		return c.JSON(http.StatusBadRequest,
			response.AsError().
				WithMessage(err.Error()))
	}

	hashedPassword, err := pkg.HashPassword(payload.Password)
	if err != nil {
		s.log.Errorln("Failed to hash password :", err)
		return c.JSON(http.StatusInternalServerError,
			response.AsError().
				WithMessage("Internal Server Error"))
	}

	payload.Password = hashedPassword

	err = s.usersRepo.Register(c.Request().Context(), payload)
	if err != nil {
		s.log.Errorln("Failed to register user :", err)
		return c.JSON(http.StatusInternalServerError,
			response.AsError().
				WithMessage("Internal Server Error"))
	}

	s.log.Println("Register Success")
	return c.JSON(http.StatusOK,
		response.WithMessage("Register Success"),
	)
}
