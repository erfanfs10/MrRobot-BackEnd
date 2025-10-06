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

func UserGetORCreate(c echo.Context) error {
	getORCreate := new(models.UserGetORCreate)
	if err := c.Bind(getORCreate); err != nil {
		return utils.HandleError(c, http.StatusBadRequest, err, "bad request")
	}
	var getORCreatedUserID int

	response := struct {
		UserID int `json:"user_id"`
	}{}

	err := db.DB.Get(&getORCreatedUserID, queries.UserGet, getORCreate.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			// create user if not exists
			err := db.DB.Get(&getORCreatedUserID, queries.UserCreate,
				getORCreate.Name, getORCreate.Email)

			if err != nil {
				return utils.HandleError(c, http.StatusInternalServerError,
					err, "server error")
			}

			// create the wishlist for user
			_, err = db.DB.Exec(queries.CreateWishList, getORCreatedUserID)
			if err != nil {
				return utils.HandleError(c, http.StatusInternalServerError, err, "server error")
			}

			response.UserID = getORCreatedUserID
			return c.JSON(http.StatusCreated, response)

		} else {
			return utils.HandleError(c, http.StatusInternalServerError,
				err, "server error")
		}
	}

	// update last login
	_, err = db.DB.Exec(queries.UpdateLastLogin, getORCreatedUserID)
	if err != nil {
		errText := fmt.Sprintf("can not update last login for user %v", getORCreatedUserID)
		c.Set("err", errText)
	}
	response.UserID = getORCreatedUserID
	return c.JSON(http.StatusOK, response)

}
