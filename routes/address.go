package routes

import (
	"github.com/erfanfs10/MrRobot-BackEnd/handlers"
	"github.com/erfanfs10/MrRobot-BackEnd/middlewares"
	"github.com/labstack/echo/v4"
)

func AddressRoutes(g *echo.Group) {
	g.GET(":id/", handlers.AddressGet, middlewares.Authenticate())
	// get current user addresses
	g.GET("user/", handlers.AddressUser, middlewares.Authenticate())
	g.POST("", handlers.AddressCreate, middlewares.Authenticate())
	g.PATCH(":id/", handlers.AddressUpdate, middlewares.Authenticate())
	g.DELETE(":id/", handlers.AddressDelete, middlewares.Authenticate())
}
