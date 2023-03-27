package main

import "fmt"

func first() []int {
	var res []int
	for i := 0; i <= 100; i++ {
		res = append(res, i)
	}
	return res
}

func second(data []int) {
	for i := 0; i <= len(data); i++ {
		fmt.Println(data[i])
	}
}

func main() {
	data := first()
	second(data)
}
