package routes

import (
	"github.com/erfanfs10/MrRobot-BackEnd/handlers"
	"github.com/labstack/echo/v4"
)

func SearchRoute(g *echo.Group) {
	g.GET("", handlers.Search)
}
