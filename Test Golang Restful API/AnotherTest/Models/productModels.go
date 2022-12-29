package models

import (
	connection "AnotherTest/Connection"
	helper "AnotherTest/Helper"
	"context"
	"net/http"
)

type Product struct {
	Id          int    `json:"id" bun:",pk"`
	Category    string `json:"category"`
	ProductName string `json:"product_name"`
	Weight      string `json:"weight"`
	Price       int    `json:"price"`
}

var err error
var res helper.Response
var db = connection.GetDB()	
var ctx = context.Background()

func GetAllProduct() (helper.Response, error) {

	products := make([]Product, 0)

	err = db.NewSelect().Model((&products)).Scan(ctx)
	if err != nil {
		return res, err
	}
	res.Data = products
	return res, nil
}

func DeleteProduct(id int) (helper.Response, error) {
	result, err := db.NewDelete().Model((*Product)(nil)).Where("id = ?", id).Returning("id").Exec(ctx, &id)
	if err != nil {
		return res, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "SUKSES DELETE Product"
	res.Data = map[string]int64{
		"rowaffected": rowAffected,
	}

	return res, nil
}

func AddNewProduct(a Product) (helper.Response, error) {

	product := &Product{Id: a.Id, Category: a.Category, ProductName: a.ProductName, Weight: a.Weight, Price: a.Price}

	_, err := db.NewInsert().Model(product).Exec(ctx)
	if err != nil {
		return res, err
	}
	res.Status = http.StatusCreated
	res.Message = "sukses"
	res.Data = map[string]int{}
	return res, nil
}

func UpateProduct(a Product) (helper.Response, error) {
	product := new(Product)
	result, err := db.NewUpdate().Model(product).Set("price = ?", a.Price).Where("id = ?", a.Id).Exec(ctx)
	if err != nil {
		return res, err
	}
	res.Status = http.StatusCreated
	res.Message = "sukses"
	res.Data = result
	return res, nil
}

// On("CONFLICT (id) DO UPDATE").
