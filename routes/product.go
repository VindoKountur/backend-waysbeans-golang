package routes

import (
	"backendwaysbeans/handlers"
	"backendwaysbeans/pkg/middleware"
	"backendwaysbeans/pkg/mysql"
	"backendwaysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Group) {
	productRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(productRepository)

	e.GET("/products", h.FindProducts)
	e.GET("/product/:id", h.GetProduct)
	e.POST("/product", middleware.Auth(middleware.IsAdmin(middleware.UploadFile(h.CreateProduct))))
	e.DELETE("/product/:id", middleware.Auth(middleware.IsAdmin(h.DeleteProduct)))
	e.PATCH("/product/:id", middleware.Auth(middleware.IsAdmin(middleware.UploadFile(h.UpdateProduct))))
}
