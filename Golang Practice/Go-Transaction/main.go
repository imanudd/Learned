package main

import (
	"database/sql"
	"fmt"
	"golang-transaction/entity"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345678"
	dbname   = "enigmacamp"
)

var psqlinfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
	studentEnrollment := entity.StudentEnrollment{Id: 4, Student_Id: 9, Subject: "Data Analityc", Credit: 3}
	enrollstudent(studentEnrollment)
}

func enrollstudent(studentEnrollment entity.StudentEnrollment) {
	db := ConnectDB()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	insertstudentenrollment(studentEnrollment, tx)

	takencredit := getsomecreditofstudent(studentEnrollment.Student_Id, tx)
	updatestudent(takencredit, studentEnrollment.Student_Id, tx)
	err = tx.Commit()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("TRANSACTION COMMITED")
	}
}

func insertstudentenrollment(studenEnrollment entity.StudentEnrollment, tx *sql.Tx) {
	insertstudentenrollment := "INSERT INTO tx_student_enrollment(id,student_id,subject,credit) VALUES ($1,$2,$3,$4);"
	_, err := tx.Exec(insertstudentenrollment, studenEnrollment.Id, studenEnrollment.Student_Id, studenEnrollment.Subject, studenEnrollment.Credit)
	if err != nil {
		panic(err)
	}
	validate(err, "insert", tx)
}

func getsomecreditofstudent(id int, tx *sql.Tx) int {
	somecredit := "SELECT SUM(credit) FROM tx_student_enrollment WHERE student_id = $1;"
	takencredit := 0
	err := tx.QueryRow(somecredit, id).Scan(&takencredit)
	validate(err, "Select", tx)
	return takencredit
}

func updatestudent(takencredit int, studentId int, tx *sql.Tx) {
	updatestudent := "UPDATE mst_student SET taken_credit = $1 WHERE id = $2;"
	_, err := tx.Exec(updatestudent, takencredit, studentId)

	validate(err, "Update", tx)
}

func validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "Transaction ROllback !")
	} else {
		fmt.Println("Successfully " + message + " Data!")
	}
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
