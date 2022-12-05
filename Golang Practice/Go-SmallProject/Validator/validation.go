package validator

import (
	"fmt"
	"os"
)

func Inputstringvalidate(variable string) {
	if variable == "" {
		fmt.Println("DATA TIDAK BOLEH KOSONG !! ")
		os.Exit(0)
	}
}

func Inputintvalidate(variable string) {
	if len(variable) >= 13 {
		fmt.Println("Nomer Tlp Terlalu Banyak ! ")
		os.Exit(0)
	}
}

func Errorcheck(variable error, info string) {
	if variable != nil {
		panic(variable.Error())
	} else {
		fmt.Println(info + " Sukses !")
	}
}

func Checkproseserror(variable error) {
	if variable != nil {
		fmt.Println("Terjadi Kesalahan Input !")
		os.Exit(0)
	}
}
