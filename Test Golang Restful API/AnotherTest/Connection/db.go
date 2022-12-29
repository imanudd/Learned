package connection

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func GetDB() *bun.DB {
	dsn := "postgres://postgres:12345678@localhost:5432/testbun?sslmode=disable"
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		panic(err)
	}
	config.PreferSimpleProtocol = true
	sqlDb := stdlib.OpenDB(*config)
	db := bun.NewDB(sqlDb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true), // log everything
	))
	return db
}
