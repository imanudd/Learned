package product

import (
	"context"
	"net/http"
)

type ProductServiceImpl struct {
	repo ProductRepository
}

func NewProductServiceImpl(repo ProductRepository) ProductService {
	return &ProductServiceImpl{
		repo: repo,
	}
}

func (svc *ProductServiceImpl) GetAllProduct(c context.Context) (res Response, err error) {
	// TODO
	res, err = svc.repo.GetAllProduct(c)
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Successfully Get All Data"
	return res, nil

}
func (svc *ProductServiceImpl) DeleteProduct(c context.Context, id int) (res Response, err error) {
	// TODO
	res, err = svc.repo.DeleteProduct(c, id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusAccepted
	res.Message = "Succsesfully Delete Data"
	return res, nil
}
func (svc *ProductServiceImpl) AddNewProduct(c context.Context, request Product) (res Response, err error) {
	// TODO
	res, err = svc.repo.AddNewProduct(c, request)
	if err != nil {
		return res, err
	}
	res.Status = http.StatusCreated
	res.Message = "Succsessfully Add Data"
	return res, nil
}
func (svc *ProductServiceImpl) UpdateProduct(c context.Context, request Product) (res Response, err error) {
	// TODO
	res, err = svc.repo.UpateProduct(c, request)
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Successfully Update Data"
	return res, nil
}
