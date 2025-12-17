package middlewares

import (
	"strings"

	"github.com/erfanfs10/MrRobot-BackEnd/utils"
	"github.com/labstack/echo/v4"
)

func Authenticate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			var userID string

			isProduction := utils.GetEnv("IS_PRODUCTION")
			if isProduction == "yes" {
				auth := c.Request().Header.Get("Authorization")
				userID = strings.TrimPrefix(auth, "Bearer ")
			} else {
				userID = c.Request().Header.Get("user_id")
			}
			
			if userID == "" {
				return echo.ErrUnauthorized
			}
			c.Set("userID", userID)

			return next(c)
		}
	}
}
