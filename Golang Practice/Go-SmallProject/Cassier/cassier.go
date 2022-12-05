package Cassier

import (
	connection "EnigmaGoLaundry/Connection"
	models "EnigmaGoLaundry/Models"
	validator "EnigmaGoLaundry/Validator"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var menucasssier int

func CasssierMenus() int {
	fmt.Println("========================= PILIH MENU ========================")
	fmt.Println("1. Daftar Kasir ")
	fmt.Println("2. Tambah Kasir ")
	fmt.Println("3. Edit Kasir ")
	fmt.Println("4. Hapus Kasir")
	fmt.Println("5. Kembali Ke Main Menu")
	fmt.Println("6. Exit")

	fmt.Print("Pilih Menu Diatas = ")
	scanner.Scan()
	menucasssier, _ = strconv.Atoi(scanner.Text())
	return menucasssier
}

func ScanCassiers(rows *sql.Rows) []models.Cassiers {
	cassiers := []models.Cassiers{}
	var err error

	for rows.Next() {
		cassier := models.Cassiers{}
		err := rows.Scan(&cassier.Id, &cassier.Name)
		validator.Checkproseserror(err)
		cassiers = append(cassiers, cassier)
	}

	err = rows.Err()
	validator.Checkproseserror(err)
	return cassiers
}

func GetAllCassier() {
	db := connection.ConnectDB()
	defer db.Close()
	fmt.Println("======================   DAFTAR Kasir  ======================")
	sqlstatement := "SELECT * FROM cassiers order by id ASC;"
	rows, err := db.Query(sqlstatement)
	validator.Checkproseserror(err)
	defer rows.Close()

	cassiers := ScanCassiers(rows)
	fmt.Println("ID || Cassier Name ")
	for _, viewcasssier := range cassiers {
		fmt.Println("==============================================")
		fmt.Println(viewcasssier.Id, " || ", viewcasssier.Name, " || ")
	}
}

func Inputcassier() models.Cassiers {

	fmt.Print("Masukan Nama Kasir = ")
	scanner.Scan()
	cassier_nama := scanner.Text()
	validator.Inputstringvalidate(cassier_nama)

	Cassier := models.Cassiers{Name: cassier_nama}
	return Cassier
}
func AddCassier(cassier models.Cassiers) {
	db := connection.ConnectDB()
	defer db.Close()
	var err error

	sqlstatement := "INSERT INTO cassiers ( name ) VALUES ($1)"
	_, err = db.Exec(sqlstatement, cassier.Name)
	validator.Errorcheck(err, "Tambah Data Kasir ")
}
func DeleteCassier(id int) {
	db := connection.ConnectDB()
	defer db.Close()
	var err error
	sqlstatement := "DELETE FROM cassiers WHERE id=$1"
	_, err = db.Exec(sqlstatement, id)
	validator.Errorcheck(err, "Delete data kasir ")
}
func Inputid() int {
	fmt.Print("Masukan Id Cassier = ")
	scanner.Scan()
	cassier_id, _ := strconv.Atoi(scanner.Text())
	return cassier_id
}

func EditCassier(id int, cassier models.Cassiers) {
	db := connection.ConnectDB()
	defer db.Close()
	var err error
	sqlstatement := "UPDATE cassiers SET name =$2 WHERE id=$1"
	_, err = db.Exec(sqlstatement, id, cassier.Name)
	validator.Errorcheck(err, "Update Data Kasir")
}

func GetCassierByID(id int) string {
	db := connection.ConnectDB()
	defer db.Close()
	var err error

	sqlstatement := "SELECT name FROM cassiers WHERE id = $1"
	cassier := models.Cassiers{}
	err = db.QueryRow(sqlstatement, id).Scan(&cassier.Name)
	validator.Checkproseserror(err)
	var Cassier string
	Cassier = cassier.Name
	return Cassier
}
