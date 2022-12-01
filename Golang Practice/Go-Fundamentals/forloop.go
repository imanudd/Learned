package main

import "fmt"

func main() {
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for i := 1; i <= 3; i++ {
	// 	fmt.Println(i)
	// }

	// i := 5
	// for {
	// 	fmt.Printf("loop ke %d \n", i)
	// 	if i == 0 {
	// 		break
	// 	}
	// 	i--
	// 	time.Sleep(1 * time.Second)
	// }

	// for n := 0; n <= 10; n++ {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Printf("bilangan Ganjil %d \n", n)
	// }

	var input int
	fmt.Scanln(&input)

	for i := 0; input > i; input-- {
		fmt.Printf("%d I will become a go developer\n", input)
	}
}
