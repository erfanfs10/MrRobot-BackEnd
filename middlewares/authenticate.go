package middlewares

import (
	"github.com/labstack/echo/v4"
)

func Authenticate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := c.Request().Header.Get("user_id")
			if userID == "" {
				return echo.ErrUnauthorized
			}
			c.Set("userID", userID)

			return next(c)
		}
	}
}
