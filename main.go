package main

import (
	"github.com/erfanfs10/MrRobot-BackEnd/db"
	"github.com/erfanfs10/MrRobot-BackEnd/middlewares"
	"github.com/erfanfs10/MrRobot-BackEnd/routes"
	"github.com/erfanfs10/MrRobot-BackEnd/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"strings"
)

func init() {
	utils.LoadEnv()
	db.ConnectToDB()
}

func main() {
	e := echo.New()
	e.Validator = utils.CreateCustomValidator()

	e.Use(middlewares.CustomLogger())
	e.Use(middlewares.SeparateLogs())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(utils.GetEnv("ALLOWED_CORS"), ","),
		AllowMethods: []string{echo.GET, echo.POST,
			echo.PUT, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin,
			echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAuthorization},
		AllowCredentials: true, // This is critical for allowing cookies and credentials
	}))

	if os.Args[len(os.Args)-1] == "dev" {
		e.Static("static", "../")
	}

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "MrRobot BackEnd")
	})

	routes.BrandRoutes(e.Group("api/brands/"))
	routes.CategoryRoutes(e.Group("api/categories/"))
	routes.ProductTypeRoutes(e.Group("api/product-types/"))
	routes.ProductRoutes(e.Group("api/products/"))
	routes.WishListRoutes(e.Group("api/wishlists/"))
	routes.AddressRoutes(e.Group("api/addresses/"))
	routes.PostRoutes(e.Group("api/posts/"))
	routes.UserRoutes(e.Group("api/users/"))
	routes.FilterRoutes(e.Group("api/filters/"))
	routes.SearchRoute(e.Group("api/search/"))
	routes.OrderRoutes(e.Group("api/orders/"))

	e.Logger.Fatal(e.Start(":8080"))
}
