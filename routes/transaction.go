package routes

import (
	"backendwaysbeans/handlers"
	"backendwaysbeans/pkg/middleware"
	"backendwaysbeans/pkg/mysql"
	"backendwaysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository, cartRepository)

	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	e.GET("/transactions", h.FindTransaction)
	e.GET("/transaction/:id", h.GetTransaction)
}
