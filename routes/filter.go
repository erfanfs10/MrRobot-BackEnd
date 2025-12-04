package routes

import (
	"github.com/erfanfs10/MrRobot-BackEnd/handlers"
	"github.com/labstack/echo/v4"
)

func FilterRoutes(g *echo.Group) {
	g.GET(":productTypeTitle/", handlers.GetFiltersProductType)
}
