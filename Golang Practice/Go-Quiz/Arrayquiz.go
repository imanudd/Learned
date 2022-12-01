package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var kapasitas int
	fmt.Scanln(&kapasitas)

	inputelement := bufio.NewScanner(os.Stdin)
	inputelement.Scan()
	element := inputelement.Text()
	var Array = strings.Split(element, " ")
	for a := 0; a < kapasitas; a++ {
		convert, err := strconv.Atoi(Array[a])
		if err == nil {
			if convert%2 == 0 {
				fmt.Println(convert)
			}
		}
	}
}

// Doni ingin mengetahui ada berapa angka genap yang terdapat di sebuah array, hanya saja doni butuh bantuan untuk mengetahuinya karena elemen di dalam array tersebut sangat banyak. Tulislah code untuk menerima 2 inputan dari Doni,
// - Inputan pertama adalah kapasitas dari array
// - Inputan kedua adalah elemen dari array tersebut.
// Lalu cetak angka-angka genap tersebut ke dalam console.
