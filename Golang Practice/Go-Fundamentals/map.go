package main

import "fmt"

func main() {

	// initiate the map, int for index or key and strings for value of index
	bulan := map[int]string{
		1: "jan",
		2: "Feb",
		3: "Mar",
		4: "April",
	}
	// to adding some value
	bulan[5] = "Mei"
	fmt.Println(bulan)

	// this another way to initiate map with "MAKE"
	var sales = make(map[string]int)
	sales["Jan"] = 2
	sales["feb"] = 4
	sales["Mar"] = 5

	// For Range to show or fetch data from map sales
	fmt.Println(sales)
	for i, j := range sales {
		fmt.Printf("Sales Bulan %s , Sejumlah %d\n", i, j)
	}

	// This method to check map of sales is exsist or not, Error Or Not
	var value, isExsist = sales["feb"]
	if isExsist {
		fmt.Println(value)
	} else {
		fmt.Println("Data Tidak Ditemukan")
	}
}
