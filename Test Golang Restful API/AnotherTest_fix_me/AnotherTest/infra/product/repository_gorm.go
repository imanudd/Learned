package product

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ProductGormRepositoryImpl struct {
	db *gorm.DB
}

func NewProductGormRepositoryImpl(connectionStr string) ProductRepository {
	// TODO
	// initialize gorm db here
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	return &ProductGormRepositoryImpl{
		db: db,
	}
}

func (repo *ProductGormRepositoryImpl) GetAllProduct(ctx context.Context) (res Response, err error) {
	// TODO
	product := []Product{}
	result := repo.db.Find(&product)
	if result.Error != nil {
		panic("Gagal Show All Product")
	}
	res.Status = http.StatusOK
	res.Message = "Berhasil Query Select Product"
	res.Data = product
	return res, nil
}

func (repo *ProductGormRepositoryImpl) DeleteProduct(ctx context.Context, id int) (res Response, err error) {
	// TODO
	var product Product
	result := repo.db.Where("id = ?", id).Delete(&product)
	if result.Error != nil {
		panic("Error Delete Data")
	}
	res.Status = http.StatusAccepted
	res.Message = "Succsesfully Delete Data"
	res.Data = map[string]int{
		"Id Product = ": id,
	}
	return res, nil
}

func (repo *ProductGormRepositoryImpl) AddNewProduct(ctx context.Context, req Product) (res Response, err error) {
	// TODO

	result := repo.db.Create(&req)
	if result.Error != nil {
		panic("Error Insert Data")
	}
	res.Status = http.StatusCreated
	res.Message = "Succsessfully Create Data"
	res.Data = req
	return res, nil
}

func (repo *ProductGormRepositoryImpl) UpateProduct(ctx context.Context, req Product) (res Response, err error) {
	// TODO
	product := []Product{}
	result := repo.db.Model(&product).Where("id = ?", req.Id).Update("price", &req.Price)
	if result.Error != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Successfully Update Data"
	res.Data = product
	return res, nil
}
