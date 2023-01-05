package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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

func AddCoin(c echo.Context) error {

	koin := c.FormValue("koin")
	nama := c.FormValue("nama")

	intKoin, err := strconv.ParseInt(koin, 0, 64)

	result, err := models.AddCoin(nama, intKoin)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}
