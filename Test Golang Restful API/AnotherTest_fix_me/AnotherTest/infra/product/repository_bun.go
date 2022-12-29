package product

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type ProductBunRepositoryImpl struct {
	db *bun.DB
}

func NewProductBunRepositoryImpl(connectionStr string) ProductRepository {
	// TODO
	// initialize db here
	fmt.Println(connectionStr)
	config, err := pgx.ParseConfig(connectionStr)
	if err != nil {
		fmt.Println("DB ERROR")
	}
	config.PreferSimpleProtocol = true
	sqlDb := stdlib.OpenDB(*config)
	db := bun.NewDB(sqlDb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	return &ProductBunRepositoryImpl{
		db: db,
	}
}

func (repo *ProductBunRepositoryImpl) GetAllProduct(ctx context.Context) (res Response, err error) {
	// TODO
	products := make([]Product, 0)

	err = repo.db.NewSelect().Model((&products)).Scan(ctx)
	if err != nil {
		return res, err
	}
	res.Data = products
	return res, nil
}

func (repo *ProductBunRepositoryImpl) DeleteProduct(ctx context.Context, id int) (res Response, err error) {
	// TODO
	result, err := repo.db.NewDelete().Model((*Product)(nil)).Where("id = ?", id).Returning("id").Exec(ctx, &id)
	if err != nil {
		return res, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Data = map[string]int64{
		"rowaffected": rowAffected,
	}

	return res, nil
}

func (repo *ProductBunRepositoryImpl) AddNewProduct(ctx context.Context, req Product) (res Response, err error) {
	// TODO
	product := &Product{Id: req.Id, Category: req.Category, ProductName: req.ProductName, Weight: req.Weight, Price: req.Price}

	result, err := repo.db.NewInsert().Model(product).Exec(ctx)
	if err != nil {
		return res, err
	}
	res.Data = map[string]interface{}{
		"new insert": result,
	}

	return res, nil
}

func (repo *ProductBunRepositoryImpl) UpateProduct(ctx context.Context, req Product) (res Response, err error) {
	// TODO
	product := new(Product)
	result, err := repo.db.NewUpdate().Model(product).Set("price = ?", req.Price).Where("id = ?", req.Id).Exec(ctx)
	if err != nil {
		return res, err
	}
	res.Data = result
	return res, nil
}
