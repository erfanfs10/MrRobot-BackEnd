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

func ProductListProductType(c echo.Context) error {
	product_type_title := c.Param("title")
	products := []models.Product{}
	query := queries.BuildProductPTCBQuery("product_type")
	err := db.DB.Select(&products, query, product_type_title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "product not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}
	return c.JSON(http.StatusOK, products)
}

func ProductListBrand(c echo.Context) error {
	brand_title := c.Param("title")
	products := []models.Product{}
	query := queries.BuildProductPTCBQuery("brand")
	err := db.DB.Select(&products, query, brand_title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "product not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}
	return c.JSON(http.StatusOK, products)
}

func ProductListCategory(c echo.Context) error {
	category_title := c.Param("title")
	products := []models.Product{}
	query := queries.BuildProductPTCBQuery("category")
	err := db.DB.Select(&products, query, category_title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "product not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}
	return c.JSON(http.StatusOK, products)
}

func ProductListNew(c echo.Context) error {
	products := []models.Product{}
	query := queries.BuildProductListQuery("new")
	err := db.DB.Select(&products, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "product not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}
	return c.JSON(http.StatusOK, products)
}

func ProductListView(c echo.Context) error {
	products := []models.Product{}
	query := queries.BuildProductListQuery("view")
	err := db.DB.Select(&products, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "product not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}
	return c.JSON(http.StatusOK, products)
}

func ProductListSell(c echo.Context) error {
	products := []models.Product{}
	query := queries.BuildProductListQuery("sell")
	err := db.DB.Select(&products, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "product not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}
	return c.JSON(http.StatusOK, products)
}

func ProductDetail(c echo.Context) error {
	product_title := c.Param("title")

	product := models.ProductDetail{}
	rates := []models.Rates{}
	images := []models.ProductImages{}
	attributes := []models.Attributes{}
	wishListItemsProductIDs := []int{}

	err := db.DB.Get(&product, queries.ProductDetail, product_title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "product not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	err = db.DB.Select(&rates, queries.Rates, *product.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "product not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	err = db.DB.Select(&images, queries.ProductImages, *product.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "product not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	err = db.DB.Select(&attributes, queries.Attributes, *product.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "product not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	product.Rates = rates
	product.Images = images
	product.Attributes = attributes
	product.WishListsProductIDs = wishListItemsProductIDs

	_, err = db.DB.Exec(queries.ProductUpdateView, product_title)
	if err != nil {
		errText := fmt.Sprintf("can not update view for product : %v", product_title)
		c.Set("err", errText)
	}
	return c.JSON(http.StatusOK, product)
}
