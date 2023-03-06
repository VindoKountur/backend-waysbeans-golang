package routes

import (
	"backendwaysbeans/handlers"
	"backendwaysbeans/pkg/middleware"
	"backendwaysbeans/pkg/mysql"
	"backendwaysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func ProfileRoutes(e *echo.Group) {
	profileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(profileRepository)

	e.PATCH("/profile", middleware.Auth(middleware.UploadFile(h.UpdateProfileByUser)))
}
