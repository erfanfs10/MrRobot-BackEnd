package routes

import (
	"github.com/erfanfs10/MrRobot-BackEnd/handlers"
	"github.com/labstack/echo/v4"
)

func BrandRoutes(g *echo.Group) {
	g.GET("", handlers.BrandList)
	g.GET(":title/", handlers.BrandGet)

}
