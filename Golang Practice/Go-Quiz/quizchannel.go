package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	channel := make(chan int)
	defer close(channel)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukan Angka = ")
	scanner.Scan()
	angka, _ := strconv.Atoi(scanner.Text())
	wg.Add(2)
	go Tentukan(channel, angka)
	go Output(channel)
	wg.Wait()
}

func Tentukan(channel chan int, a int) {
	defer wg.Done()
	if a%3 == 0 {
		channel <- a
	} else {
		channel <- 0
	}
}

func Output(channel chan int) {
	defer wg.Done()
	if <-channel == 0 {
		fmt.Println("Bukan Angka Ganjil")
	} else {
		fmt.Println("Angka Ganjil")
	}

}
