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
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository, cartRepository, productRepository)

	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	e.GET("/transactions-user", middleware.Auth(h.GetUserTransaction))
	e.GET("/transactions", h.FindTransaction)
	e.GET("/transaction/:id", h.GetTransaction)
	e.POST("/notification", h.Notification)
}
