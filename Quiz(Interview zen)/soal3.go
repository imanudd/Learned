package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// I create a array for save url links
	var url [4]string
	url[0] = "https://id.bitdegree.org"
	url[1] = "https://finance.detik.com"
	url[2] = "https://telkom.co.id"
	url[3] = "https://corona.jakarta.go.id"

	simpan := Loop(*Regex(), url[:])
	Condition(simpan, url[:])

}

// this func for compile regular expression
func Regex() *regexp.Regexp {
	var regex, _ = regexp.Compile(`([-a-z0-9]+)(?:\.co.id|.com|.go.id|.org)`)
	return regex
	//return must be = detik.com etc..
}

// this func for loop result after manipulate url or show url
func Loop(regex regexp.Regexp, url []string) []string {
	var simpan []string
	for i := 0; i < len(url); i++ {
		str := strings.Split(regex.Split(url[i], -1)[0], "https://")
		simpan = append(simpan, str[1])
	}
	return simpan
	// this return must be subdomain of url like "corona." etc ...
}

// this func use for check condition if sub domain is exist or not!
func Condition(simpan []string, url []string) {
	for i := 0; i < len(simpan); i++ {
		if simpan[i] == "" {
			fmt.Println("Url : ", url[i], "Tidak Memiliki Sub Domain")
		} else {
			fmt.Println("Url : ", url[i], "Memiliki Sub Domain = ", simpan[i])
		}
	}

}
