package routes

import (
	"github.com/erfanfs10/MrRobot-BackEnd/handlers"
	"github.com/erfanfs10/MrRobot-BackEnd/middlewares"
	"github.com/labstack/echo/v4"
)

func OrderRoutes(g *echo.Group) {
	g.POST("", handlers.OrderCreate, middlewares.Authenticate())
	g.GET("my/", handlers.OrderList, middlewares.Authenticate())
}
