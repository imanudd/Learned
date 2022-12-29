package routes

import (
	controller "AnotherTest/Controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", hello)
	e.GET("/products", controller.GetAllProduct)
	e.DELETE("/products", controller.DeleteProduct)
	e.POST("/products", controller.AddNewProduct)
	e.PATCH("/products/:id", controller.UpdateProduct)
	// e.GET("/viewproduct", controller.GetAllProduct)

	return e

}
