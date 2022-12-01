package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	var inputan = []string{}
	for i := 1; i <= 3; i++ {
		input.Scan()
		inputan = append(inputan, input.Text())
	}

	fmt.Printf("Halo, saya %s. Saya tinggal di %s. Alamat email saya adalah %s", inputan[0], inputan[1], inputan[2])

}
