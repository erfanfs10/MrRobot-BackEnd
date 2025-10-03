package routes

import (
	"github.com/erfanfs10/MrRobot-BackEnd/handlers"
	"github.com/labstack/echo/v4"
)

func CategoryRoutes(g *echo.Group) {
	g.GET("", handlers.CategoryList)
	g.GET(":title/", handlers.CategoryGet)
}
