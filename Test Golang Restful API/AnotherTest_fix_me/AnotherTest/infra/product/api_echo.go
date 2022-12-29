package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo, svc *API) {
	e.GET("/products", svc.GetAllProduct)
	e.DELETE("/products/:id", svc.DeleteProduct)
	e.POST("/products", svc.AddNewProduct)
	e.PATCH("/products/:id", svc.UpdateProduct)
	e.GET("/hello", func(c echo.Context) error { fmt.Println("helo"); return nil })
	fmt.Println("check regist")
}

type API struct {
	svc ProductService
}

func NewProductAPIImpl(svc ProductService) *API {
	return &API{
		svc: svc,
	}
}

func (api *API) GetAllProduct(c echo.Context) error {
	// TODO
	res, err := api.svc.GetAllProduct(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, res)
}

func (api *API) DeleteProduct(c echo.Context) error {
	// TODO
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := api.svc.DeleteProduct(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, res)
}

func (api *API) AddNewProduct(c echo.Context) error {
	// TODO
	var product Product
	product.Id, _ = strconv.Atoi(c.FormValue("product_id"))
	product.Category = c.FormValue("category")
	product.ProductName = c.FormValue("product_name")
	product.Weight = c.FormValue("weight")
	product.Price, _ = strconv.Atoi(c.FormValue("price"))

	res, err := api.svc.AddNewProduct(c.Request().Context(), product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}

func (api *API) UpdateProduct(c echo.Context) error {
	// TODO
	var product Product
	product.Id, _ = strconv.Atoi(c.Param("id"))
	product.Price, _ = strconv.Atoi(c.FormValue("price"))
	res, err := api.svc.UpdateProduct(c.Request().Context(), product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}

func (api *API) Hello() error {
	fmt.Println("hello world")
	return nil
}
