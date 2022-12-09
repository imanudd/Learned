package models

import (
	"GO-QuizAPI/db"
	"GO-QuizAPI/helper"
	"database/sql"
	"fmt"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.Init()

	sqlStatement := "SELECT id,username,password FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(&obj.Id, &obj.Username, &pwd)

	if err == sql.ErrNoRows {
		fmt.Println("Username Not Found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query ERROR")
		return false, err
	}

	match, err := helper.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash And Password Doesn't Match !")
		return false, err
	}
	return true, nil
}
