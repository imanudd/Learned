package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	text := []int{1, 4, 6, 10, 14}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukan Angka nya = ")
	scanner.Scan()
	var min []int
	var max []int
	input, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < len(text); i++ {
		if text[i]-input < 0 {
			min = append(min, text[i]-input)
			bubblesort(min)
		} else if text[i]-input > 0 {
			max = append(max, text[i]-input)
			bubblesort(max)
		}
	}
	fmt.Println("Angka Yang Paling Dekat Kebawah = ", min[len(min)-1]+input)
	fmt.Println("Angka Yang Paling Dekat Keatas = ", max[0]+input)
}

func bubblesort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		isSort := true
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				isSort = false
			}
		}

		if isSort {
			break
		}

	}
	return array
}
