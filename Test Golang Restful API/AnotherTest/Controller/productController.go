package controller

import (
	models "AnotherTest/Models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllProduct(c echo.Context) error {
	res, err := models.GetAllProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, res)
}

func DeleteProduct(c echo.Context) error {
	product_id, _ := strconv.Atoi(c.FormValue("product_id"))
	result, err := models.DeleteProduct(product_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusNoContent, result)
}

func AddNewProduct(c echo.Context) error {
	var product models.Product
	product.Id, _ = strconv.Atoi(c.FormValue("id"))
	product.Category = c.FormValue("category")
	product.ProductName = c.FormValue("product_name")
	product.Weight = c.FormValue("weight")
	product.Price, _ = strconv.Atoi(c.FormValue("price"))
	result, err := models.AddNewProduct(product)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}
func UpdateProduct(c echo.Context) error {
	var product models.Product
	product.Id, _ = strconv.Atoi(c.Param("id"))

	product.Price, _ = strconv.Atoi(c.FormValue("price"))
	result, err := models.UpateProduct(product)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}
