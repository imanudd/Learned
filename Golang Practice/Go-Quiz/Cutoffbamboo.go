package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	jmlbambo, _ := strconv.Atoi(scanner.Text())
	sekat := []int{}

	for i := 1; i <= jmlbambo; i++ {
		scanner.Scan()
		input, _ := strconv.Atoi(scanner.Text())
		sekat = append(sekat, input)
	}
	scanner.Scan()
	cut, _ := strconv.Atoi(scanner.Text())
	for k := 0; k < cut; k++ {
		fmt.Printf("Potongan ke- %d\n", k+1)
		for f := 0; f < len(sekat); f++ {
			if sekat[f] == 0 {
				sekat[f] = sekat[f]
				fmt.Println(sekat[f])
			} else {
				sekat[f] = sekat[f] - 1
				fmt.Println(sekat[f])
			}
		}
	}
}
