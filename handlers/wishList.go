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
	"github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func WishListUser(c echo.Context) error {
	userID := c.Get("userID")
	var wishListID int
	wishListItemsProductIDs := []int{}
	wishListItems := []models.Product{}

	// get the wishlist id from db for current user
	err := db.DB.Get(&wishListID, queries.GetWishListIDCurrentUser, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "wishlist not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	// get the product id's of wishlist items for current wishlist id
	err = db.DB.Select(&wishListItemsProductIDs, queries.GetWishListItemsProductIDs, wishListID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "product not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	if len(wishListItemsProductIDs) == 0 {
		return c.JSON(http.StatusOK, wishListItems)
	}

	// create the query for IN
	query, args, err := sqlx.In(queries.GetWishListItems, wishListItemsProductIDs)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	// rebind every thing and return actual query
	query = db.DB.Rebind(query)

	// get current user wish list items from db wish Product model
	err = db.DB.Select(&wishListItems, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "product not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	return c.JSON(http.StatusOK, wishListItems)
}

func WishListCreate(c echo.Context) error {

	userID := c.Get("userID")
	productID := new(models.ProductID)
	var wishListID int
	wishListItemsProductIDs := []int{}

	if err := c.Bind(productID); err != nil {
		return utils.HandleError(c, http.StatusBadRequest, err, "bad request")
	}

	// get the wishlist id from db for current user
	err := db.DB.Get(&wishListID, queries.GetWishListIDCurrentUser, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "wishlist not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	_, err = db.DB.Exec(queries.CreateWishListItem, wishListID, productID.ProductID)
	if err != nil {
		// var pgErr *pgconn.PgError
		// if errors.As(err, &pgErr) {
		// 	if pgErr.Code == "23505" {
		// 		return utils.HandleError(c, http.StatusConflict, err, "already added")
		// 	}
		// }
		errText := fmt.Sprintf("product id %v for wishlist id %v for user id %v already added : dupicate", productID, wishListID, userID)
		c.Set("err", errText)

	}

	// get the product id's of wishlist items for current wishlist id
	err = db.DB.Select(&wishListItemsProductIDs, queries.GetWishListItemsProductIDs, wishListID)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err, "server error")
	}

	return c.JSON(http.StatusCreated, wishListItemsProductIDs)
}

func WishListDelete(c echo.Context) error {
	userID := c.Get("userID")
	productID := c.Param("id")
	var wishListID int
	wishListItemsProductIDs := []int{}

	// get the wishlist id from db for current user
	err := db.DB.Get(&wishListID, queries.GetWishListIDCurrentUser, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "wishlist not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	_, err = db.DB.Exec(queries.DeleteWishListItem, wishListID, productID)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err, "server error")
	}

	// get the product id's of wishlist items for current wishlist id
	err = db.DB.Select(&wishListItemsProductIDs, queries.GetWishListItemsProductIDs, wishListID)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err, "server error")
	}

	return c.JSON(http.StatusOK, wishListItemsProductIDs)
}
