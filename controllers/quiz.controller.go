package controllers

import (
	"net/http"
	"strconv"
	"vp_giver_echo/models"

	"github.com/labstack/echo/v4"
)

func FetchAllQuiz(c echo.Context) error {

	result, err := models.FetchAllQuiz()
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchQuizSubbab(c echo.Context) error {

	subbab := c.Param("subbab")
	intsubbab, err := strconv.ParseInt(subbab, 0, 64)

	result, err := models.FetchQuizSubbab(intsubbab)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreQuiz(c echo.Context) error {

	pertanyaan := c.FormValue("pertanyaan")
	jawaban := c.FormValue("jawaban")
	subbab := c.FormValue("subbab")
	intsubbab, err := strconv.ParseInt(subbab, 0, 64)

	result, err := models.StoreQuiz(pertanyaan, jawaban, intsubbab)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CheckQuiz(c echo.Context) error {
	pertanyaan := c.FormValue("pertanyaan")
	jawaban := c.FormValue("jawaban")

	res, err := models.CheckQuiz(pertanyaan, jawaban)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message1": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message2": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Jawaban Benar",
	})
}
