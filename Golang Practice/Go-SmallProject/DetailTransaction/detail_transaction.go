package detailtransaction

import (
	"EnigmaGoLaundry/Cassier"
	connection "EnigmaGoLaundry/Connection"
	"EnigmaGoLaundry/Customers"
	models "EnigmaGoLaundry/Models"
	services "EnigmaGoLaundry/Services"
	transaction "EnigmaGoLaundry/Transaction"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func ScanDetailTransactions(rows *sql.Rows) []models.Detail_Transaction {
	detailtransactions := []models.Detail_Transaction{}
	var err error

	for rows.Next() {
		detailtransaction := models.Detail_Transaction{}

		err := rows.Scan(&detailtransaction.Id, &detailtransaction.Idtransaction, &detailtransaction.IdService, &detailtransaction.Qty)
		if err != nil {
			panic(err)
		}
		detailtransactions = append(detailtransactions, detailtransaction)

	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return detailtransactions
}
func GetAllDetailTransaction() {
	db := connection.ConnectDB()
	defer db.Close()
	fmt.Println("====================== DAFTAR TRANSAKSI ======================")
	sqlstatement := "select * from detail_transactions;"
	rows, err := db.Query(sqlstatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	detailtransactions := ScanDetailTransactions(rows)
	for _, viewdetailtransaction := range detailtransactions {
		fmt.Println("==========================================================")
		fmt.Println(viewdetailtransaction.Id, " || ", viewdetailtransaction.Idtransaction, " || ", viewdetailtransaction.IdService, " || ", viewdetailtransaction.Qty, "||")
	}
}

func InputDetailTransaction() models.Detail_Transaction {

	fmt.Print("Masukan id service = ")
	scanner.Scan()
	service_id, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Masukan Jumlah (qty) = ")
	scanner.Scan()
	id := transaction.GetTheLastID().Id
	qty, _ := strconv.Atoi(scanner.Text())
	Detail_Transaction := models.Detail_Transaction{Idtransaction: id, IdService: service_id, Qty: qty}
	return Detail_Transaction

}

func JumlahPelayanan() {
	fmt.Print("Masukan Jumlah Pelayanan = ")
	scanner.Scan()
	service, _ := strconv.Atoi(scanner.Text())
	services.GetAllService()
	for i := 1; i <= service; i++ {
		AddDetailTransaction(InputDetailTransaction())
	}
}

func AddDetailTransaction(detail models.Detail_Transaction) {
	db := connection.ConnectDB()
	defer db.Close()
	var err error

	sqlstatement := "INSERT INTO detail_transaction (id_transaction,id_service ,qty) VALUES ($1,$2,$3);"
	_, err = db.Exec(sqlstatement, detail.Idtransaction, detail.IdService, detail.Qty)
	if err != nil {
		panic(err)
	}

}
func GetDetailByID(id int) {
	db := connection.ConnectDB()
	defer db.Close()
	var err error

	sqlstatement := "SELECT * FROM detail_transaction WHERE id_transaction = $1"
	rows, err := db.Query(sqlstatement, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var hasil int
	detailtransactions := ScanDetailTransactions(rows)
	var i = 1
	for _, viewdetailtransaction := range detailtransactions {
		fmt.Println("====================================================================================")
		fmt.Println(i, "||", services.GetServiceByID(viewdetailtransaction.IdService).ServiceName, " || ", viewdetailtransaction.Qty, "||", services.GetServiceByID(viewdetailtransaction.IdService).Unit, "||", services.GetServiceByID(viewdetailtransaction.IdService).Price, "||", services.GetServiceByID(viewdetailtransaction.IdService).Price*viewdetailtransaction.Qty)
		hasil += services.GetServiceByID(viewdetailtransaction.IdService).Price * viewdetailtransaction.Qty
		i += 1
	}
	fmt.Println("====================================================================================")
	fmt.Printf("================================   GRAND TOTAL  =  %d ===============================", hasil)
	fmt.Println()
	fmt.Println("====================================================================================")
	fmt.Println()
}

func ViewInvoice(id int) {
	cust := Customers.GetCustomerByID(transaction.GetTransactionByID(id).IdCustomer)
	cassier := Cassier.GetCassierByID(transaction.GetTransactionByID(id).IdCassier)
	infotrans := transaction.GetTransactionByID(id)
	fmt.Println("================================== ENIGMA LAUNDRY ==================================")

	fmt.Printf("NO TRANSAKSI : %d                                      Nama Cust : %s", id, cust.Name)
	fmt.Println()
	fmt.Printf("Tanggal Masuk : %s                    No Hp :  %s", infotrans.StartDate, cust.Phone_number)
	fmt.Println()
	fmt.Printf("Tanggal Selesai : %s ", infotrans.EndDate)
	fmt.Println()
	fmt.Printf("Diterima Oleh : %s ", cassier)
	fmt.Println()
	fmt.Println("====================================================================================")
	fmt.Println("No   || Pelayanan        || Jumlah        || Satuan       || Harga        || Total  ")
	fmt.Println()
	GetDetailByID(id)
	fmt.Println("==================================   THANK YOU    ==================================")

}
