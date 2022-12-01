package main

import "fmt"

func main() {
	tipeproduct := "dryfood"

	switch tipeproduct {
	case "dryfood":
		fmt.Println("True")
	case "wetfood":
		fmt.Println("not dry food is wet food")
	default:
		fmt.Println("ini Nilai Default")
	}

	var nilai = 80
	switch {
	case nilai > 80:
		fmt.Println("Nilai A")
	case nilai > 70:
		fmt.Println("Nilai B")
	default:
		fmt.Println("Nilali C")
	}
}
