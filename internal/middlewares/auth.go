package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/wildanfaz/e-ticket-terminal/internal/helpers"
	"github.com/wildanfaz/e-ticket-terminal/internal/pkg"
	"github.com/wildanfaz/e-ticket-terminal/internal/repositories"
)

func Auth(usersRepo repositories.UsersRepository, log *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			response := helpers.NewResponse()
			bearerToken := c.Request().Header.Get("Authorization")
			token := ""

			if len(strings.Split(bearerToken, " ")) > 1 {
				token = strings.Split(bearerToken, " ")[1]
			}

			claims, err := pkg.ValidateToken(token)
			if err != nil {
				log.Errorln("Failed to validate token", err)
				return c.JSON(http.StatusUnauthorized,
					response.AsError().
						WithMessage("Unauthorized"),
				)
			}

			user, err := usersRepo.GetUserByEmail(c.Request().Context(), claims.Email)
			if err != nil {
				log.Errorln("Failed to get user", err)
				return c.JSON(http.StatusUnauthorized,
					response.AsError().
						WithMessage("Unauthorized"),
				)
			}

			if user == nil {
				log.Errorln("User not found")
				return c.JSON(http.StatusUnauthorized,
					response.AsError().
						WithMessage("Unauthorized"),
				)
			}

			c.Set("email", claims.Email)

			return next(c)
		}
	}
}
