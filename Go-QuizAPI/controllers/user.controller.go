package controllers

import (
	"GO-QuizAPI/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAllData(c echo.Context) error {
	result, err := models.FetchAllData()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}
