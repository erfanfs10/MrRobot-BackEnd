package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/erfanfs10/MrRobot-BackEnd/db"
	"github.com/erfanfs10/MrRobot-BackEnd/models"
	"github.com/erfanfs10/MrRobot-BackEnd/queries"
	"github.com/erfanfs10/MrRobot-BackEnd/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func OrderCreate(c echo.Context) error {
	userID := c.Get("userID")

	orderCreate := new(models.OrderCreate)
	if err := c.Bind(orderCreate); err != nil {
		return utils.HandleError(c, http.StatusBadRequest, err, "bad request")
	}

	var newOrderID int

	trackingNumber := uuid.New().String()

	err := db.DB.Get(&newOrderID, queries.OrderCreate,
		userID, orderCreate.AddressID, orderCreate.TotalProducts,
		orderCreate.ShippingPrice, "payed", trackingNumber)
	if err != nil {
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	for _,v := range(*orderCreate.Cart) {
		_, err := db.DB.Exec(queries.OrderItemsCreate, newOrderID, v.ID, v.Quantity, v.NetPrice)
		if err != nil {
			return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "created"})

}

func OrderList(c echo.Context) error {
	userID := c.Get("userID")

	orders := []models.Order{}

	err := db.DB.Select(&orders, queries.OrderList, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.HandleError(c, http.StatusNotFound, err, "orders not found")
		}
		return utils.HandleError(c, http.StatusInternalServerError,
			err, "server error")
	}

	return c.JSON(http.StatusOK, orders)
}
