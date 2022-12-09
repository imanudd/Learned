package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

const (
	host     = "127.0.0.1"
	port     = "3306"
	user     = "root"
	password = ""
	dbname   = "echorest"
)

func Init() *sql.DB {
	sqlStatement := fmt.Sprintf("%s:@tcp(%s:%s)/%s", user, host, port, dbname)
	db, err := sql.Open("mysql", sqlStatement)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
