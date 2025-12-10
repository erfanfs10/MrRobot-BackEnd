package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/erfanfs10/MrRobot-BackEnd/db"
	"github.com/erfanfs10/MrRobot-BackEnd/models"
	"github.com/erfanfs10/MrRobot-BackEnd/queries"
	"github.com/erfanfs10/MrRobot-BackEnd/utils"
	"github.com/labstack/echo/v4"
)

func Search(c echo.Context) error {
	q := c.QueryParam("q")
	fmt.Println(q)

	products := []models.Product{}
	err := db.DB.Select(&products, queries.SearchProduct, "%"+q+"%")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "product not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}
	return c.JSON(http.StatusOK, products)
}
