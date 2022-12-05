package transaction

import (
	"EnigmaGoLaundry/Cassier"
	connection "EnigmaGoLaundry/Connection"
	"EnigmaGoLaundry/Customers"
	models "EnigmaGoLaundry/Models"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var menutransaction int

func TransactionMenus() int {
	fmt.Println("========================= PILIH MENU ========================")
	fmt.Println("1. Daftar Transaksi ")
	fmt.Println("2. Tambah Transaksi Baru ")
	fmt.Println("3. Kembali ke Main Menu")
	fmt.Println("4. Keluar ")

	fmt.Print("Pilih Menu Diatas = ")
	scanner.Scan()
	menutransaction, _ = strconv.Atoi(scanner.Text())
	return menutransaction
}
func ScanTransactions(rows *sql.Rows) []models.Transaction {
	transactions := []models.Transaction{}
	var err error

	for rows.Next() {
		transaction := models.Transaction{}

		err := rows.Scan(&transaction.Id, &transaction.IdCustomer, &transaction.IdCassier, &transaction.StartDate, &transaction.EndDate)
		if err != nil {
			panic(err)
		}
		transactions = append(transactions, transaction)

	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return transactions
}
func GetAllTransaction() {
	db := connection.ConnectDB()
	defer db.Close()
	fmt.Println("====================== DAFTAR TRANSAKSI ======================")
	sqlstatement := "select * from transactions;"
	rows, err := db.Query(sqlstatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	transactions := ScanTransactions(rows)
	fmt.Println("ID || Customer Name || Cassier Name || Start Date || End Date")
	for _, viewtransaction := range transactions {
		fmt.Println("==========================================================")
		fmt.Println(viewtransaction.Id, " || ", Customers.GetCustomerByID(viewtransaction.IdCustomer).Name, " || ", Cassier.GetCassierByID(viewtransaction.IdCassier), " || ", viewtransaction.StartDate, "||", viewtransaction.EndDate)
	}

}

func GetTransactionByID(id int) models.Transaction {
	db := connection.ConnectDB()
	defer db.Close()
	var err error

	transaction := models.Transaction{}
	sqlstatement := "SELECT * FROM transactions where id = $1;"
	err = db.QueryRow(sqlstatement, id).Scan(&transaction.Id, &transaction.IdCustomer, &transaction.IdCassier, &transaction.StartDate, &transaction.EndDate)
	if err != nil {
		panic(err)
	}

	return transaction
}

func InputTransaction() models.Transaction {
	Customers.GetAllCustomer()
	fmt.Print("Masukan Id Customer = ")
	scanner.Scan()
	customers_id, _ := strconv.Atoi(scanner.Text())
	Cassier.GetAllCassier()
	fmt.Print("Masukan Id Kasir = ")
	scanner.Scan()
	cassier_id, _ := strconv.Atoi(scanner.Text())
	fmt.Print("Masukan Tanggal Masuk (yyyy-mm-dd) = ")
	scanner.Scan()
	Start_date := scanner.Text()
	fmt.Print("Masukan Tanggal Keluar (yyyy-mm-dd) = ")
	scanner.Scan()
	end_date := scanner.Text()

	Transaction := models.Transaction{IdCustomer: customers_id, IdCassier: cassier_id, StartDate: Start_date, EndDate: end_date}
	return Transaction

}

func AddTransaction(transaction models.Transaction) {
	db := connection.ConnectDB()
	defer db.Close()
	var err error

	sqlstatement := "INSERT INTO transactions (id_customer,id_cassier ,start_date,end_date) VALUES ($1,$2,$3,$4);"
	_, err = db.Exec(sqlstatement, transaction.IdCustomer, transaction.IdCassier, transaction.StartDate, transaction.EndDate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Sukses")
	}
}

func GetTheLastID() models.Transaction {
	db := connection.ConnectDB()
	defer db.Close()
	var err error

	sqlstatement := "SELECT * FROM transactions order by id desc limit 1"
	transaction := models.Transaction{}
	err = db.QueryRow(sqlstatement).Scan(&transaction.Id, &transaction.IdCustomer, &transaction.IdCassier, &transaction.StartDate, &transaction.EndDate)
	if err != nil {
		panic(err)
	}

	return transaction
}
