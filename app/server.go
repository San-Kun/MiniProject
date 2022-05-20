package app

import (
	"fmt"

	"github.com/San-Kun/MiniProject/app/config"
	"github.com/San-Kun/MiniProject/docs"
	_bankAccountController "github.com/San-Kun/MiniProject/user/delivery/http"
	_bankController "github.com/San-Kun/MiniProject/user/delivery/http"
	_eventController "github.com/San-Kun/MiniProject/user/delivery/http"
	_eventTicketController "github.com/San-Kun/MiniProject/user/delivery/http"
	_paymentController "github.com/San-Kun/MiniProject/user/delivery/http"
	_paymentTypeController "github.com/San-Kun/MiniProject/user/delivery/http"
	_userController "github.com/San-Kun/MiniProject/user/delivery/http"
	mid "github.com/San-Kun/MiniProject/user/delivery/http/middleware"
	"github.com/San-Kun/MiniProject/user/repository"
	"github.com/San-Kun/MiniProject/user/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Event App API Documentation
// @description Find an event that interests you
// @version 2.0
// @host localhost:8080
// @BasePath
// @schemes http https
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
func Run() {

	db := config.InitDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	eventRepository := repository.NewEventRepository(db)
	eventUsecase := usecase.NewEventUsecase(eventRepository)
	paymentRepository := repository.NewPaymentRepository(db)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepository)
	bankAccountRepository := repository.NewBankAccountRepository(db)
	bankAccountUsecase := usecase.NewBankAccountUsecase(bankAccountRepository)
	bankRepository := repository.NewBankRepository(db)
	bankUsecase := usecase.NewBankUsecase(bankRepository)
	paymentTypeRepository := repository.NewPaymentTypeRepository(db)
	paymentTypeUsecase := usecase.NewPaymentTypeUsecase(paymentTypeRepository)
	eventTicketRepository := repository.NewEventTicketRepository(db)
	eventTicketUsecase := usecase.NewEventTicketUsecase(eventTicketRepository)

	e := echo.New()
	mid.NewGoMiddleware().LogMiddleware(e)
	_userController.NewUserController(e, userUsecase)
	_eventController.NewEventController(e, eventUsecase)
	_paymentController.NewPaymentController(e, paymentUsecase)
	_bankAccountController.NewBankAccountController(e, bankAccountUsecase)
	_bankController.NewBankController(e, bankUsecase)
	_paymentTypeController.NewPaymentTypeController(e, paymentTypeUsecase)
	_eventTicketController.NewEventTicketController(e, eventTicketUsecase)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Event App API Documentation"
	docs.SwaggerInfo.Description = "Find an event that interests you."
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	address := fmt.Sprintf("localhost:%d", 8080)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	if err := e.Start(address); err != nil {
		log.Info("Exit The Server")
	}
}
