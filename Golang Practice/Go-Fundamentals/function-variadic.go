package main

import "fmt"

func variadicSumAll(angka ...int) int {
	total := len(angka)
	return total

}

func main() {
	angka := variadicSumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10)
	fmt.Println(angka)
}

// func main() {
// 	angka := variadicSumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10)
// 	fmt.Println(angka)
// }
