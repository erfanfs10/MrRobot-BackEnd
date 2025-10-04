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

func AddressGet(c echo.Context) error {
	userID := c.Get("userID")
	addressID := c.Param("id")

	address := models.Address{}
	err := db.DB.Get(&address, queries.AddressGet, addressID, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "address not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}
	return c.JSON(http.StatusOK, address)
}

func AddressUser(c echo.Context) error {
	userID := c.Get("userID")
	userAddresses := []models.Address{}
	err := db.DB.Select(&userAddresses, queries.AddressUser, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "user not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}
	return c.JSON(http.StatusOK, userAddresses)
}

func AddressCreate(c echo.Context) error {
	userID := c.Get("userID")
	addressCreate := new(models.AddressCreate)
	if err := c.Bind(addressCreate); err != nil {
		return utils.HandleError(c, http.StatusBadRequest, err, "bad request")
	}
	err := c.Validate(addressCreate)
	if err != nil {
		return utils.CustomValidationError(c, http.StatusBadRequest,
			err, "validation error")
	}
	var userAddressesCount uint8
	err = db.DB.Get(&userAddressesCount, queries.AddressUserCount, userID)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	if userAddressesCount <= 4 {
		newAddress := models.Address{}
		var newAddressID int
		err = db.DB.Get(&newAddressID, queries.AddressCreate,
			userID, addressCreate.Address, addressCreate.Title)
		if err != nil {
			return utils.HandleError(c, http.StatusInternalServerError,
				err, "server error")
		}

		err := db.DB.Get(&newAddress, queries.AddressGet, newAddressID, userID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return utils.HandleError(c, http.StatusNotFound, err, "address not found")
			}
			return utils.HandleError(c, http.StatusInternalServerError,
				err, "server error")
		}

		return c.JSON(http.StatusOK, newAddress)

	} else {
		err := fmt.Sprintf("can not create new address more than 5 for user : %v", userID)
		return utils.HandleError(c, http.StatusBadRequest,
			errors.New(err), "you can not add more addresses")
	}
}

func AddressUpdate(c echo.Context) error {
	userID := c.Get("userID")
	addressID := c.Param("id")
	addressUpdate := new(models.AddressUpdate)
	if err := c.Bind(addressUpdate); err != nil {
		return utils.HandleError(c, http.StatusBadRequest, err, "bad request")
	}
	err := c.Validate(addressUpdate)
	if err != nil {
		return utils.CustomValidationError(c, http.StatusBadRequest,
			err, "validation error")
	}
	result, err := db.DB.Exec(queries.AddressUpdate,
		addressUpdate.Address, addressUpdate.Title, addressID, userID)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err, "server error")
	}
	rowAffected, _ := result.RowsAffected()
	if rowAffected == 0 {
		err := fmt.Sprintf(
			"can not update address fobidden or not found user : %v, address : %v",
			userID, addressID)
		return utils.HandleError(c, http.StatusBadRequest, errors.New(err), "address not found")
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "address updated"})
}

func AddressDelete(c echo.Context) error {
	userID := c.Get("userID")
	addressID := c.Param("id")

	_, err := db.DB.Exec(queries.AddressDelete, addressID, userID)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError, err, "server error")
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "address deleted"})
}
