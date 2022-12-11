package controllers

import (
	"net/http"
	"strconv"
	"vp_giver_echo/models"

	"github.com/labstack/echo/v4"
)

func FetchAllBuku(c echo.Context) error {

	result, err := models.FetchAllBuku()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}


func FetchBukuPljrn(c echo.Context) error {

	pelajaran := c.Param("pelajaran")
	intpelajaran, err := strconv.ParseInt(pelajaran, 0, 64)

	result, err := models.FetchBukuPljrn(intpelajaran)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreBuku(c echo.Context) error {

	nama_buku := c.FormValue("nama_buku")
	pelajaran := c.FormValue("pelajaran_id")
	intpelajaran, err := strconv.ParseInt(pelajaran, 0, 64)
	imageCover := c.FormValue("imageCover")
	imageBanner := c.FormValue("imageBanner")


	result, err := models.StoreBuku(nama_buku, intpelajaran, imageCover, imageBanner)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
