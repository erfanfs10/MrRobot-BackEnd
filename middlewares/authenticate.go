package middlewares

import (
	"fmt"
	"strings"

	"github.com/erfanfs10/MrRobot-BackEnd/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Authenticate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := c.Request().Header.Get("user_id")
			fmt.Println(userID, "userID")
			auth := c.Request().Header.Get("Authorization")
			tokenStr := strings.TrimPrefix(auth, "Bearer ")
			fmt.Println(auth, "auth")
			fmt.Println(tokenStr, "tokenStr")
			sess := c.Request().Header.Get("session")
			fmt.Println(sess, "sss")
			sec := utils.GetEnv("DB_USER")
			token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				return []byte(sec), nil
			})

			if err != nil || !token.Valid {
				fmt.Println(err)
			}

			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims, "cccc")

			if userID == "" {
				return echo.ErrUnauthorized
			}
			c.Set("userID", userID)

			return next(c)
		}
	}
}
