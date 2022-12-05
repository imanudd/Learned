package services

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
var menuservice int

func ServiceMenus() int {
	fmt.Println("========================= PILIH MENU ========================")
	fmt.Println("1. Daftar Pelayanan ")
	fmt.Println("2. Tambah Pelayanan ")
	fmt.Println("3. Edit Pelayanan ")
	fmt.Println("4. Hapus Pelayanan")
	fmt.Println("5. Kembali Main Menu")
	fmt.Println("6. Keluar")

	fmt.Print("Pilih Menu Diatas = ")
	scanner.Scan()
	menuservice, _ = strconv.Atoi(scanner.Text())
	return menuservice
}
func ScanServices(rows *sql.Rows) []models.Services {
	services := []models.Services{}
	var err error

	for rows.Next() {
		service := models.Services{}
		err := rows.Scan(&service.Id, &service.ServiceName, &service.Unit, &service.Price)
		validator.Checkproseserror(err)
		services = append(services, service)
	}

	err = rows.Err()
	validator.Checkproseserror(err)
	return services
}

func GetAllService() {
	db := connection.ConnectDB()
	defer db.Close()
	fmt.Println("======================== DAFTAR PELAYANAN ======================")
	sqlstatement := "SELECT * FROM services order by id ASC;"
	rows, err := db.Query(sqlstatement)
	validator.Checkproseserror(err)
	defer rows.Close()

	services := ScanServices(rows)
	fmt.Println("ID || Service Name || Unit || Price ")
	for _, viewservice := range services {
		fmt.Println("==============================================")
		fmt.Println(viewservice.Id, " || ", viewservice.ServiceName, " || ", viewservice.Unit, " || ", viewservice.Price, " || ")
	}
}
func InputService() models.Services {

	fmt.Print("Masukan Nama Pelayanan = ")
	scanner.Scan()
	service_nama := scanner.Text()
	validator.Inputstringvalidate(service_nama)

	fmt.Print("Masukan Satuan Pelayanan = ")
	scanner.Scan()
	service_unit := scanner.Text()
	validator.Inputstringvalidate(service_unit)

	fmt.Print("Masukan Harga Pelayanan = ")
	scanner.Scan()
	service_price, err := strconv.Atoi(scanner.Text())
	validator.Checkproseserror(err)

	Services := models.Services{ServiceName: service_nama, Unit: service_unit, Price: service_price}
	return Services
}
func AddService(service models.Services) {
	db := connection.ConnectDB()
	defer db.Close()
	var err error

	sqlstatement := "INSERT INTO services (service_name ,unit, price) VALUES ($1,$2,$3)"
	_, err = db.Exec(sqlstatement, service.ServiceName, service.Unit, service.Price)
	validator.Errorcheck(err, "Tambah data layanan")
}
func Inputid() int {
	fmt.Print("Masukan Id Service = ")
	scanner.Scan()
	service_id, _ := strconv.Atoi(scanner.Text())
	return service_id
}

func EditService(id int, service models.Services) {
	db := connection.ConnectDB()
	defer db.Close()
	var err error
	sqlstatement := "UPDATE services SET service_name =$2, unit=$3,price = $4 WHERE id=$1"
	_, err = db.Exec(sqlstatement, id, service.ServiceName, service.Unit, service.Price)
	validator.Errorcheck(err, "Update Data Service")
}

func DeleteService(id int) {
	db := connection.ConnectDB()
	defer db.Close()
	var err error
	sqlstatement := "DELETE FROM services WHERE id=$1"
	_, err = db.Exec(sqlstatement, id)
	validator.Errorcheck(err, "Delete Data Service")
}

func GetServiceByID(id int) models.Services {
	db := connection.ConnectDB()
	defer db.Close()
	var err error

	sqlstatement := "SELECT service_name,unit,price FROM services WHERE id = $1"
	service := models.Services{}
	err = db.QueryRow(sqlstatement, id).Scan(&service.ServiceName, &service.Unit, &service.Price)
	validator.Checkproseserror(err)

	Services := models.Services{ServiceName: service.ServiceName, Unit: service.Unit, Price: service.Price}
	return Services

}
