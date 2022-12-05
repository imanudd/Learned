package main

import (
	"EnigmaGoLaundry/Cassier"
	"EnigmaGoLaundry/Customers"
	detailtransaction "EnigmaGoLaundry/DetailTransaction"
	services "EnigmaGoLaundry/Services"
	transaction "EnigmaGoLaundry/Transaction"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)
var menuselected int
var menucustomer int
var menucasssier int
var menuservice int
var menutransaction int

func header() int {
	fmt.Println("==============================================================================")
	fmt.Println("                           WELCOME TO ENIGMA LAUNDRY                          ")
	fmt.Println("================================== MAIN MENU =================================")
	fmt.Println(" 1. Customers || 2. Kasir || 3. Daftar Pelayanan || 4. Transaksi || 5. Exit   ")
	fmt.Println("==============================================================================")
	fmt.Print("PILIH MENU YANG TERSEDIA : ")
	scanner.Scan()
	menuselected, _ = strconv.Atoi(scanner.Text())
	return menuselected
}
func main() {
	for {
		header()
		if menuselected > 4 {
			break
		} else if menuselected == 1 {
			for {
				menucustomer = Customers.CustomerMenus()
				if menucustomer >= 6 {
					os.Exit(0)
				} else if menucustomer == 1 {
					Customers.GetAllCustomer()
				} else if menucustomer == 2 {
					Customers.AddCustomer(Customers.Inputcustomer())
				} else if menucustomer == 3 {
					Customers.GetAllCustomer()
					fmt.Println("======================  EDIT CUSTOMER  ======================")
					Customers.EditCustomer(Customers.Inputid(), Customers.Inputcustomer())
				} else if menucustomer == 4 {
					Customers.GetAllCustomer()
					Customers.DeleteCustomer(Customers.Inputid())
				} else if menucustomer == 5 {
					break
				}
			}
		} else if menuselected == 2 {
			for {
				menucasssier = Cassier.CasssierMenus()
				if menucasssier > 5 {
					os.Exit(0)
				} else if menucasssier == 1 {
					Cassier.GetAllCassier()
				} else if menucasssier == 2 {
					Cassier.AddCassier(Cassier.Inputcassier())
				} else if menucasssier == 3 {
					Cassier.GetAllCassier()
					fmt.Println("======================  EDIT KASIR  ======================")
					Cassier.EditCassier(Cassier.Inputid(), Cassier.Inputcassier())
				} else if menucasssier == 4 {
					Cassier.GetAllCassier()
					Cassier.DeleteCassier(Cassier.Inputid())
				} else if menucasssier == 5 {
					break
				}
			}
		} else if menuselected == 3 {
			for {
				menuservice = services.ServiceMenus()
				if menuservice > 5 {
					os.Exit(0)
				} else if menuservice == 1 {
					services.GetAllService()
				} else if menuservice == 2 {
					services.AddService(services.InputService())
				} else if menuservice == 3 {
					services.GetAllService()
					fmt.Println("==================== EDIT PELAYANAN ====================")
					services.EditService(services.Inputid(), services.InputService())
				} else if menuservice == 4 {
					services.GetAllService()
					services.DeleteService(services.Inputid())
				} else if menuservice == 5 {
					break
				}
			}

		} else if menuselected == 4 {
			for {
				menutransaction = transaction.TransactionMenus()
				if menutransaction == 1 {
					transaction.GetAllTransaction()
				} else if menutransaction == 2 {
					transaction.AddTransaction(transaction.InputTransaction())
					detailtransaction.JumlahPelayanan()
					detailtransaction.ViewInvoice(transaction.GetTheLastID().Id)
				} else if menutransaction == 3 {
					break
				} else if menutransaction == 4 {
					os.Exit(0)
				}

			}
		}
	}
}
