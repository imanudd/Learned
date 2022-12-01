package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "Herman budi jenny kevin aris"
	slice := strings.Split(text, " ")
	//fmt.Println(slice)
	hasil := make([]string, len(slice))
	for i := 0; i < len(slice); i++ {
		if len(slice[i])%2 == 0 {
			hasil[i] = slice[i]
		}
	}

	fmt.Println(hasil)

}
