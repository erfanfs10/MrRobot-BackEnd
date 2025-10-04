package routes

import (
	"github.com/erfanfs10/MrRobot-BackEnd/handlers"
	"github.com/labstack/echo/v4"
)

func PostRoutes(g *echo.Group) {
	g.GET("category/:title/", handlers.PostListCategory)
	g.GET("tag/:title/", handlers.PostListTag)
	g.GET(":slug/", handlers.PostDetail)
	g.GET("random/", handlers.PostRandom)
}
