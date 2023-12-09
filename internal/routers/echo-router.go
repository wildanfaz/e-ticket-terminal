package routers

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/wildanfaz/e-ticket-terminal/configs"
	"github.com/wildanfaz/e-ticket-terminal/internal/constants"
	"github.com/wildanfaz/e-ticket-terminal/internal/middlewares"
	"github.com/wildanfaz/e-ticket-terminal/internal/pkg"
	"github.com/wildanfaz/e-ticket-terminal/internal/repositories"
	"github.com/wildanfaz/e-ticket-terminal/internal/services/health"
	"github.com/wildanfaz/e-ticket-terminal/internal/services/terminals"
	"github.com/wildanfaz/e-ticket-terminal/internal/services/users"
)

func InitEchoRouter() {
	fmt.Println(constants.Blue, "---Init Echo Router---")

	// configs
	config := configs.InitConfig()
	dbMySql := configs.InitMySql()

	// pkg
	log := pkg.InitLogger()

	// repositories
	usersRepo := repositories.NewUsersRepository(dbMySql)
	terminalsRepo := repositories.NewTerminalsRepository(dbMySql)

	// services
	usersService := users.NewService(usersRepo, log)
	terminalsService := terminals.NewService(terminalsRepo, usersRepo, log)

	e := echo.New()

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/health", health.HealthCheck)

	// middlewares
	auth := middlewares.Auth(usersRepo, log)

	// users
	users := apiV1.Group("/users")
	users.POST("/register", usersService.Register)
	users.POST("/login", usersService.Login)
	users.POST("/topup", usersService.TopUp, auth)

	// terminals
	terminals := apiV1.Group("/terminals")
	terminals.POST("", terminalsService.AddTerminal, auth)
	terminals.POST("/transaction", terminalsService.AddTransaction, auth)
	terminals.PUT("/transaction", terminalsService.UpdateTransaction, auth)

	e.Logger.Fatal(e.Start(config.EchoPort))
}
