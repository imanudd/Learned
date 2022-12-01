package main

import "fmt"

func main() {
	var myval int = 64
	var myvalpointer *int = &myval
	fmt.Println(myval)
	fmt.Println(myvalpointer)
	fmt.Println(*myvalpointer)

	var myval2 int
	fmt.Println(myval2)
}
