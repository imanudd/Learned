package main

import "fmt"

func sayHelo() {
	fmt.Println("Hello")
}

func sayHello(firstname string, lastname string) {
	fmt.Printf("Hello %s %s\n", firstname, lastname)
}

func increment(angka int) int {
	return 1 + angka
}

func main() {
	sayHelo()
	sayHello("MUhammad Imanuddin", "Jamil")
	angka := increment(1)
	fmt.Println(angka)
}
