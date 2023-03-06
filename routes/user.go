package routes

import (
	"backendwaysbeans/handlers"
	"backendwaysbeans/pkg/mysql"
	"backendwaysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerUser(userRepository, transactionRepository)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:id", h.GetUser)
}
