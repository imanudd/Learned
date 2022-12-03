package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "761165575"
	var save []string
	s := strings.Split(text, "")

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				save = append(save, s[j])
				break
			}
		}
	}
	for j := len(save) - 2; j >= 0; j-- {
		save = append(save, save[j])
	}

	convert := strings.Join(save, "")
	fmt.Print(convert)
}
