package main

import (
	"github.com/erfanfs10/MrRobot-BackEnd/db"
	"github.com/erfanfs10/MrRobot-BackEnd/middlewares"
	"github.com/erfanfs10/MrRobot-BackEnd/routes"
	"github.com/erfanfs10/MrRobot-BackEnd/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

func init() {
	utils.LoadEnv()
	db.ConnectToDB()
}

func main() {
	e := echo.New()

	e.Use(middlewares.CustomLogger())
	e.Use(middlewares.SeparateLogs())
	e.Use(middleware.Recover())

	if os.Args[len(os.Args)-1] == "dev" {
		e.Static("static", "../")
	}

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "MrRobot BackEnd")
	})

	routes.BrandRoutes(e.Group("api/brands/"))
	e.Logger.Fatal(e.Start(":8080"))
}
