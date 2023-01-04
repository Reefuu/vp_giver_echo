package controllers

import (
	"net/http"
	"strconv"
	"vp_giver_echo/models"

	"github.com/labstack/echo/v4"
)

func FetchAllPelajaran(c echo.Context) error {

	result, err := models.FetchAllPelajaran()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchPelajaranKls(c echo.Context) error {

	kelas := c.Param("kelas")
	intkelas, err := strconv.ParseInt(kelas, 0, 64)

	result, err := models.FetchPelajaranKls(intkelas)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StorePelajaran(c echo.Context) error {

	nama_pelajaran := c.FormValue("nama_pelajaran")
	kelas := c.FormValue("kelas")
	intkelas, err := strconv.ParseInt(kelas, 0, 64)
	image := c.FormValue("image")
	imageBanner := c.FormValue("imageBanner")

	result, err := models.StorePelajaran(nama_pelajaran, intkelas, image, imageBanner)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
