package main

import (
	"backendwaysbeans/database"
	"backendwaysbeans/pkg/mysql"
	"backendwaysbeans/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mysql.DatabaseInit()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	e.Static("/uploads", "./uploads")

	PORT := "5000"
	e.Logger.Fatal(e.Start("localhost:" + PORT))
}
