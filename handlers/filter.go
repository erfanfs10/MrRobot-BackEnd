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

func GetFiltersProductType(c echo.Context) error {

	pt_title := c.Param("productTypeTitle")
	fmt.Println(pt_title)

	attributeFilters := []models.AttributeFilters{}
	err := db.DB.Select(&attributeFilters, queries.GetFilterAttributes, pt_title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "attributeFilters not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	brandFilters := []models.Brand{}
	err = db.DB.Select(&brandFilters, queries.GetBrandsProductTypes, pt_title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "brand not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	attributeFilterValues := []models.AttributeFilterValue{}
	err = db.DB.Select(&attributeFilterValues, queries.GetFilterAttributeValues, pt_title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "attributeFilterValues not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	// Group values
	attrMap := map[int][]string{}
	for _, r := range attributeFilterValues {
		attrMap[r.AttributeID] = append(attrMap[r.AttributeID], r.Title)
	}

	// Attach values
	for i := range attributeFilters {
		attributeFilters[i].Values = attrMap[attributeFilters[i].ID]
	}

	filters := struct {
		Brands     []models.Brand            `json:"brands"`
		Attributes []models.AttributeFilters `json:"attributes"`
	}{}

	filters.Brands = brandFilters
	filters.Attributes = attributeFilters

	// Final Response
	return c.JSON(http.StatusOK, filters)
}
