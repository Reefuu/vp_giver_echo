package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"vp_giver_echo/models"
)

func RegisterNewAcc(c echo.Context) error {

	nama := c.FormValue("nama")
	email := c.FormValue("email")
	password := c.FormValue("password")

	result, err := models.RegisterNewAcc(nama, email, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
