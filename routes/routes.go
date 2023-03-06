package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	ProductRoutes(e)
	AuthRoutes(e)
	UserRoutes(e)
	ProfileRoutes(e)
	AddressRoutes(e)
	TransactionRoutes(e)
	CartRoutes(e)
}
