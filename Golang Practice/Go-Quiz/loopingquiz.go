package main

import "fmt"

func main() {
	var input int
	fmt.Scanln(&input)

	for i := 0; input > i; input-- {
		fmt.Printf("%d  I will become a go developer\n", input)
	}

}
