package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var x int
	var y float64

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ = strconv.ParseFloat(scanner.Text(), 32)

	hasil1 := x + int(y)
	hasil2 := float64(x) - y
	hasil3 := int(y) * x

	fmt.Print(hasil1)
	fmt.Printf("\n%.2f", hasil2)
	fmt.Print("\n", hasil3)
}

// // Pada tugas kali ini kamu akan diberikan dua variabel yang pertama dalam bentuk integer dan yang kedua dalam bentuk float32. Tulislah code yang akan menampilkan pada console berupa :
// 1. Baris pertama menampilkan penjumlahan dari dua variabel yang diberikan dalam bentuk integer
// 2. Baris kedua menampilkan pengurangan dalam bentuk float dengan 2 koma
// 3. Baris ketiga menampilkan perkalian dalam bentuk integer
