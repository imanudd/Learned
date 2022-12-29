package product

import (
	"context"
)

type ProductRepository interface {
	GetAllProduct(ctx context.Context) (res Response, err error)
	DeleteProduct(ctx context.Context, id int) (res Response, err error)
	AddNewProduct(ctx context.Context, req Product) (res Response, err error)
	UpateProduct(ctx context.Context, req Product) (res Response, err error)
}

type ProductService interface {
	GetAllProduct(c context.Context) (res Response, err error)
	DeleteProduct(c context.Context, id int) (res Response, err error)
	AddNewProduct(c context.Context, request Product) (res Response, err error)
	UpdateProduct(c context.Context, request Product) (res Response, err error)
}
