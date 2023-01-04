package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"vp_giver_echo/models"
)

func FetchAllUsers(c echo.Context) error {

	result, err := models.FetchAllUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
