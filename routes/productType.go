package routes

import (
	"github.com/erfanfs10/MrRobot-BackEnd/handlers"
	"github.com/labstack/echo/v4"
)

func ProductTypeRoutes(g *echo.Group) {
	g.GET(":category-id/", handlers.ProductTypeCategory)
	g.GET("", handlers.ProductTypeList)
}
