package models

import (
	"GO-QuizAPI/db"
	"net/http"
)

type DetailUser struct {
	Id        int    `json:"id"`
	Id_user   int    `json:"id_user"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Address   string `json:"address"`
}

func FetchAllData() (Response, error) {
	var user DetailUser
	var arruser []DetailUser
	var res Response

	con := db.Init()

	sqlStatement := "SELECT * FROM detail_user"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Id_user, &user.Firstname, &user.Lastname, &user.Address)
		if err != nil {
			return res, err
		}
		arruser = append(arruser, user)
	}
	res.Status = http.StatusOK
	res.Message = "Succsess"
	res.Data = arruser

	return res, nil
}
