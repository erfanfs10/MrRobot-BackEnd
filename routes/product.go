package routes

import (
	"github.com/erfanfs10/MrRobot-BackEnd/handlers"
	"github.com/labstack/echo/v4"
)

func ProductRoutes(g *echo.Group) {
	g.GET("product-type/:title/", handlers.ProductListProductType)
	g.GET("brand/:title/", handlers.ProductListBrand)
	g.GET("category/:title/", handlers.ProductListCategory)
	g.GET("new/", handlers.ProductListNew)
	g.GET("view/", handlers.ProductListView)
	g.GET("sell/", handlers.ProductListSell)
}
