package routes

import (
	"GO-QuizAPI/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Succsessfully Connected")
	})
	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)
	e.GET("/detailuser", controllers.FetchAllData, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("secret")}))

	return e

}
