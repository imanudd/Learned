package main

import "fmt"

func main() {
	// number := 2
	// number2 := 2

	// if number%2 == 0 && number2%2 == 1 {
	// 	fmt.Println("Nilai Ganjil Dan Genap")
	// } else if number%2 == 0 {
	// 	fmt.Println("Nilai Genap")
	// } else {
	// 	fmt.Println("Nilai Ganjil")
	// }

	// if waktu := 3; waktu >= 6 && waktu <= 10 {
	// 	fmt.Println("Breakfast")
	// } else if waktu > 10 && waktu < 14 {
	// 	fmt.Println("Lunch")
	// } else {
	// 	fmt.Println("Dinner")
	// }

	var jam int
	fmt.Scanln(&jam)

	if jam >= 24 {
		fmt.Println("Hanya ada 24 jam dalam sehari")
	} else if jam >= 4 && jam < 6 {
		fmt.Println("Bangun Pagi")
	} else if jam >= 6 && jam < 8 {
		fmt.Println("Mandi dan Sarapan")
	} else if jam >= 8 && jam < 11 {
		fmt.Println("Berangkat Sekolah")
	} else if jam == 12 {
		fmt.Println("Pulang Sekolah")
	} else if jam >= 13 && jam < 15 {
		fmt.Println("Tidur Siang")
	} else if jam >= 16 && jam < 18 {
		fmt.Println("Waktu Main")
	} else {
		fmt.Println("Waktu Belajar dan Istirahat")
	}

}
