package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type BangunDatar interface {
	Luas() int
}

type Segitiga struct {
	alas   int
	tinggi int
}

func (luas *Segitiga) Luas() int {
	return (luas.alas * luas.tinggi) / 2
}

//kode struct disini

//kode struct method disini

func getLuas(bangunDatar BangunDatar) {
	fmt.Printf("Luas bangun datar = %d \n", bangunDatar.Luas())
}

func main() {
	var segitiga1 Segitiga
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	segitiga1.alas, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	segitiga1.tinggi, _ = strconv.Atoi(scanner.Text())
	getLuas(&segitiga1)

}
