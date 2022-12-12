package controllers

import (
	"net/http"
	"strconv"
	"vp_giver_echo/models"

	"github.com/labstack/echo/v4"
)

func FetchAllSubbab(c echo.Context) error {

	result, err := models.FetchAllSubbab()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}


func FetchSubbabBab(c echo.Context) error {

	bab := c.Param("bab")
	intbab, err := strconv.ParseInt(bab, 0, 64)

	result, err := models.FetchSubbabBab(intbab)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreSubbab(c echo.Context) error {

	materi := c.FormValue("materi")
	bab := c.FormValue("bab_id")
	intbab, err := strconv.ParseInt(bab, 0, 64)


	result, err := models.StoreSubbab(materi, intbab)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
