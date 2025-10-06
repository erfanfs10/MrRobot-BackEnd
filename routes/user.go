package routes

import (
	"github.com/erfanfs10/MrRobot-BackEnd/handlers"
	"github.com/labstack/echo/v4"
)

func UserRoutes(g *echo.Group) {
	g.POST("get-or-create/", handlers.UserGetORCreate)
}
