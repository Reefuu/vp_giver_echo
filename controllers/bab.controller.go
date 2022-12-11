package controllers

import (
	"net/http"
	"strconv"
	"vp_giver_echo/models"

	"github.com/labstack/echo/v4"
)

func FetchAllBab(c echo.Context) error {

	result, err := models.FetchAllBab()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}


func FetchBabBuku(c echo.Context) error {

	buku := c.Param("buku")
	intbuku, err := strconv.ParseInt(buku, 0, 64)

	result, err := models.FetchBabBuku(intbuku)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreBab(c echo.Context) error {

	bab_nama := c.FormValue("bab_nama")
	buku := c.FormValue("buku_id")
	intbuku, err := strconv.ParseInt(buku, 0, 64)
	imageCover := c.FormValue("imageCover")
	imageBanner := c.FormValue("imageBanner")


	result, err := models.StoreBab(bab_nama, intbuku, imageCover, imageBanner)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
