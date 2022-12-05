package Customers

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
var menucustomer int

func CustomerMenus() int {
	fmt.Println("========================= PILIH MENU ========================")
	fmt.Println("1. Daftar Customer ")
	fmt.Println("2. Tambah Customer ")
	fmt.Println("3. Edit Customer ")
	fmt.Println("4. Hapus Customer")
	fmt.Println("5. Kembali ke Main Menu")
	fmt.Println("6. Keluar")

	fmt.Print("Pilih Menu Diatas = ")
	scanner.Scan()
	menucustomer, _ = strconv.Atoi(scanner.Text())
	return menucustomer
}

func ScanCustomers(rows *sql.Rows) []models.Customers {
	customers := []models.Customers{}
	var err error

	for rows.Next() {
		customer := models.Customers{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Phone_number)
		validator.Checkproseserror(err)
		customers = append(customers, customer)
	}

	err = rows.Err()
	validator.Checkproseserror(err)
	return customers
}

func GetAllCustomer() {
	db := connection.ConnectDB()
	defer db.Close()
	fmt.Println("====================== DAFTAR CUSTOMER ======================")
	sqlstatement := "SELECT * FROM customers order by id ASC;"
	rows, err := db.Query(sqlstatement)
	validator.Checkproseserror(err)
	defer rows.Close()

	customers := ScanCustomers(rows)
	fmt.Println("ID || Customer Name || Phone Number ")
	for _, viewcustomer := range customers {
		fmt.Println("==============================================")
		fmt.Printf("%2d || %8s || %01s \n", viewcustomer.Id, viewcustomer.Name, viewcustomer.Phone_number)
	}
}

func Inputcustomer() models.Customers {
	fmt.Print("Masukan Nama Customer = ")
	scanner.Scan()
	customer_nama := scanner.Text()
	validator.Inputstringvalidate(customer_nama)

	fmt.Print("Masukan Nomer Telphon Customer = ")
	scanner.Scan()
	inputph, error := strconv.Atoi(scanner.Text())
	validator.Checkproseserror(error)

	customer_ph := strconv.Itoa(inputph)
	validator.Inputintvalidate(customer_ph)

	Customers := models.Customers{Name: customer_nama, Phone_number: customer_ph}
	return Customers
}
func AddCustomer(customer models.Customers) {
	db := connection.ConnectDB()
	defer db.Close()
	var err error

	sqlstatement := "INSERT INTO customers (name ,phone_number) VALUES ($1,$2)"
	_, err = db.Exec(sqlstatement, customer.Name, customer.Phone_number)
	validator.Errorcheck(err, "Insert Data Customers")
	GetAllCustomer()
}
func DeleteCustomer(id int) {
	db := connection.ConnectDB()
	defer db.Close()
	var err error

	sqlstatement := "DELETE FROM customers WHERE id=$1"
	_, err = db.Exec(sqlstatement, id)
	validator.Errorcheck(err, "Delete Data Customer ")
}
func Inputid() int {
	fmt.Print("Masukan Id Customer = ")
	scanner.Scan()
	customer_id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return customer_id
}

func EditCustomer(id int, customer models.Customers) {
	db := connection.ConnectDB()
	defer db.Close()
	var err error

	sqlstatement := "UPDATE customers SET name =$2, phone_number=$3 WHERE id=$1"
	_, err = db.Exec(sqlstatement, id, customer.Name, customer.Phone_number)
	validator.Errorcheck(err, "Update Data Customer")
}
func GetCustomerByID(id int) models.Customers {
	db := connection.ConnectDB()
	defer db.Close()
	var err error

	sqlstatement := "SELECT * FROM customers WHERE id = $1"
	customer := models.Customers{}
	err = db.QueryRow(sqlstatement, id).Scan(&customer.Id, &customer.Name, &customer.Phone_number)
	validator.Checkproseserror(err)
	return customer
}
