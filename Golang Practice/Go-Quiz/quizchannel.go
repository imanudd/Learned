package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	channel := make(chan int)
	defer close(channel)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukan Angka = ")
	scanner.Scan()
	angka, _ := strconv.Atoi(scanner.Text())
	go Tentukan(channel, angka)
	Output(channel)
}

func Tentukan(channel chan int, a int) {
	if a%3 == 0 {
		channel <- a
		return
	}
	channel <- 0
}

func Output(channel chan int) {
	if <-channel == 0 {
		fmt.Println("Bukan Angka Ganjil")
		return
	}
	fmt.Println("Angka Ganjil")

}
