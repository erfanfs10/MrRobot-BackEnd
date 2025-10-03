package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/erfanfs10/MrRobot-BackEnd/db"
	"github.com/erfanfs10/MrRobot-BackEnd/models"
	"github.com/erfanfs10/MrRobot-BackEnd/queries"
	"github.com/erfanfs10/MrRobot-BackEnd/utils"
	"github.com/labstack/echo/v4"
)

// product types based on category
func ProductTypeCategory(c echo.Context) error {
	category_id := c.Param("category-id")
	productTypes := []models.ProductType{}
	err := db.DB.Select(&productTypes, queries.ProductTypesBasedOnCategory, category_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "product types not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}
	return c.JSON(http.StatusOK, productTypes)
}

func ProductTypeList(c echo.Context) error {
	productTypes := []models.ProductType{}
	err := db.DB.Select(&productTypes, queries.ProductTypeList)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "product types not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}
	return c.JSON(http.StatusOK, productTypes)

}
