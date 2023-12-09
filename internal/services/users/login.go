package users

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/e-ticket-terminal/internal/helpers"
	"github.com/wildanfaz/e-ticket-terminal/internal/models"
	"github.com/wildanfaz/e-ticket-terminal/internal/pkg"
)

func (s *Service) Login(c echo.Context) error {
	var (
		response = helpers.NewResponse()
		payload  models.UserLogin
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
		validation.Field(&payload.Password, validation.Required),
		validation.Field(&payload.ConfirmPassword, validation.Required),
	)
	if err != nil {
		s.log.Errorln("Invalid Payload :", err)
		return c.JSON(http.StatusBadRequest,
			response.AsError().
				WithMessage(err.Error()))
	}

	if payload.Password != payload.ConfirmPassword {
		s.log.Errorln("Password and Confirm Password not match")
		return c.JSON(http.StatusBadRequest,
			response.AsError().
				WithMessage("Password and Confirm Password not match"),
		)
	}

	user, err := s.usersRepo.GetUserByEmail(c.Request().Context(), payload.Email)
	if err != nil {
		s.log.Errorln("Failed to get user :", err)
		return c.JSON(http.StatusInternalServerError,
			response.AsError().
				WithMessage("Internal Server Error"))
	}

	if user == nil {
		s.log.Errorln("User not found")
		return c.JSON(http.StatusBadRequest,
			response.AsError().
				WithMessage("Invalid Email or Password"))
	}

	if !pkg.ComparePassword(user.Password, payload.Password) {
		s.log.Errorln("Invalid Password")
		return c.JSON(http.StatusBadRequest,
			response.AsError().
				WithMessage("Invalid Email or Password"))
	}

	ss, err := pkg.GenerateToken(user.Email)
	if err != nil {
		s.log.Errorln("Failed to generate token :", err)
		return c.JSON(http.StatusInternalServerError,
			response.AsError().
				WithMessage("Internal Server Error"))
	}

	s.log.Println("Login Success")
	return c.JSON(http.StatusOK,
		response.WithMessage("Login Success").
			WithData(pkg.NewLogin(ss)),
	)
}
