package routes

import (
	"github.com/erfanfs10/MrRobot-BackEnd/handlers"
	"github.com/erfanfs10/MrRobot-BackEnd/middlewares"
	"github.com/labstack/echo/v4"
)

func WishListRoutes(g *echo.Group) {
	// get current user wishlistitems
	g.GET("user/", handlers.WishListUser, middlewares.Authenticate())
	g.POST("", handlers.WishListCreate, middlewares.Authenticate())
	g.DELETE(":id/", handlers.WishListDelete, middlewares.Authenticate())
}
