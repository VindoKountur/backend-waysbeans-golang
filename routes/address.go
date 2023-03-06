package routes

import (
	"backendwaysbeans/handlers"
	"backendwaysbeans/pkg/middleware"
	"backendwaysbeans/pkg/mysql"
	"backendwaysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func AddressRoutes(e *echo.Group) {
	addressRepository := repositories.RepositoryAddress(mysql.DB)
	h := handlers.HandlerAddress(addressRepository)

	// User Address
	e.GET("/address", middleware.Auth(h.FindAddressByUser))
	e.POST("/address", middleware.Auth(h.CreateAddressByUser))
	e.PATCH("/address/:id", middleware.Auth(h.UpdateAddress))
	e.DELETE("/address/:id", middleware.Auth(h.DeleteAddress))
}
