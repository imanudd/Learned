package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var nama string
	// fmt.Printf("Nama Lengkap = ")
	// fmt.Scanf("%s ", &nama) //Tidak Bisa Menggunakan Spasi
	// fmt.Println(nama)

	// fmt.Printf("Alamat = ")
	// input := os.Args
	// nama := input[1]
	// alamat := input[2]
	// email := input[3]
	// fmt.Println(nama, alamat, email)
	// fmt.Println("Halo, saya", nama, ". saya tinggal di", alamat, ". Alamat email saya adalah", email)

	// 	Name := flag.String("Name", "", "User Name")
	// 	Homeaddress := flag.String("Address", "", "User Address")
	// 	email := flag.String("email", "", "User Email")

	// 	flag.Parse()

	// 	fmt.Println("Halo, saya", *Name, ". saya tinggal di", *Homeaddress, ". Alamat email saya adalah", *email)

	//
	// input := os.Args
	// nama := input[1]
	// //fmt.Println(nama)

	// alamat := input[2]
	// fmt.Println(nama, alamat)
	// inputemail := os.Args
	// email := inputemail

	// fmt.Println(nama, alamat, email)
	inputnama := bufio.NewScanner(os.Stdin)
	inputnama.Scan()
	nama := inputnama.Text()

	inputalamat := bufio.NewScanner(os.Stdin)
	inputalamat.Scan()
	alamat := inputalamat.Text()

	inputemail := bufio.NewScanner(os.Stdin)
	inputemail.Scan()
	email := inputemail.Text()
	fmt.Println("Halo, saya", nama, ". Saya tinggal di", alamat, ". Alamat email saya adalah", email)

}
