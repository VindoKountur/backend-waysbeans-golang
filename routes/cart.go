package routes

import (
	"backendwaysbeans/handlers"
	"backendwaysbeans/pkg/mysql"
	"backendwaysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func CartRoutes(e *echo.Group) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	e.GET("/carts", h.FindCarts)
}
