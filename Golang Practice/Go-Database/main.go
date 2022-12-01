package main

import (
	"GODB/entiity"
	"database/sql"
	"fmt"

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
	//AddStudent("6", "jamil", "Jamil@mail.com", "jakarta", "1998-04-20", "M")
	// student := entiity.Student{Id: "10", Name: "MARCHO", Email: "marcho@gmail.com", Address: "JEKARDAH", Birth_date: time.Date(1999, 10, 11, 0, 0, 0, 0, time.Local), Gender: "M"}
	//AddStudent(student)
	//UpdateStudent(student)
	// DeleteStudent("10")
	// students := GetAllStudent()
	// for _, student := range students {
	// 	fmt.Println(student.Id, student.Name, student.Email, student.Address, student.Birth_date, student.Gender, student.TakenCredit)
	// }

	fmt.Println(GetStudenByID(5))
	// students := SearchBy("me", "1996-02-25")
	// for _, student := range students {
	// 	fmt.Println(student.Id, student.Name, student.Email, student.Address, student.Birth_date, student.Gender)
	// }

}

func AddStudent(student entiity.Student) {
	db := ConnectDB()
	defer db.Close()
	var err error

	sqlstatement := "INSERT INTO mst_student (id, name ,email, address, birth_date, gender) VALUES ($1,$2,$3,$4,$5,$6)"
	_, err = db.Exec(sqlstatement, student.Id, student.Name, student.Email, student.Address, student.Birth_date, student.Gender)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Insert Data")
	}
}

func UpdateStudent(student entiity.Student) {
	db := ConnectDB()
	defer db.Close()
	var err error

	sqlstatement := "UPDATE mst_student SET name =$2, email=$3, address=$4, birth_date=$5, gender=$6 WHERE id=$1"
	_, err = db.Exec(sqlstatement, student.Id, student.Name, student.Email, student.Address, student.Birth_date, student.Gender)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Update Data")
	}
}

func DeleteStudent(id string) {
	db := ConnectDB()
	defer db.Close()
	var err error
	sqlstatement := "DELETE FROM mst_student WHERE id=$1"
	_, err = db.Exec(sqlstatement, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Delete Data")
	}
}

func GetAllStudent() []entiity.Student {
	db := ConnectDB()
	defer db.Close()

	sqlstatement := "SELECT * FROM mst_student;"
	rows, err := db.Query(sqlstatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	students := ScanStudent(rows)
	return students
}
func SearchBy(name string, birthdate string) []entiity.Student {
	db := ConnectDB()
	defer db.Close()
	sqlstatement := "SELECT * FROM mst_student WHERE name LIKE $1 AND birth_date = $2"
	rows, err := db.Query(sqlstatement, "%"+name+"%", birthdate)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	students := ScanStudent(rows)
	return students
}

func GetStudenByID(id int) entiity.Student {
	db := ConnectDB()
	defer db.Close()
	var err error
	students := []entiity.Student{}

	sqlstatement := "SELECT * FROM mst_student WHERE id = $1"
	student := entiity.Student{}
	err = db.QueryRow(sqlstatement, id).Scan(&student.Id, &student.Name, &student.Email, &student.Address, &student.Birth_date, &student.Gender, &student.TakenCredit)
	if err != nil {
		panic(err)
	} else {
		students = append(students, student)
	}
	return student
}

func ScanStudent(rows *sql.Rows) []entiity.Student {
	students := []entiity.Student{}
	var err error

	for rows.Next() {
		student := entiity.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Address, &student.Birth_date, &student.Gender, &student.TakenCredit)
		if err != nil {
			panic(err)
		}
		students = append(students, student)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return students
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
