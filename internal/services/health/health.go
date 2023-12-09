package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/e-ticket-terminal/internal/helpers"
)

func HealthCheck(c echo.Context) error {
	var (
		response = helpers.NewResponse()
	)

	return c.JSON(http.StatusOK, response.WithMessage("OK"))
}
