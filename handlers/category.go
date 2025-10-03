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

func CategoryList(c echo.Context) error {
	categories := []models.Category{}
	err := db.DB.Select(&categories, queries.CategoryList)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "category not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}
	return c.JSON(http.StatusOK, categories)

}

func CategoryGet(c echo.Context) error {
	category_title := c.Param("title")
	category := models.Category{}
	err := db.DB.Get(&category, queries.CategoryGet, category_title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "category not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}
	return c.JSON(http.StatusOK, category)

}
