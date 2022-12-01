package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Penduduk struct {
	nama, alamat string
	umur         int
}

type Student struct {
	// Isi struct ini
	Name string
	Age  int
}

func (s *Student) Introduction() {
	fmt.Printf("Hello, my name is %s. Im %d years old", s.Name, s.Age)
}

func main() {
	// var penduduk1 Penduduk
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// penduduk1.nama = scanner.Text()
	// scanner.Scan()
	// penduduk1.umur, _ = strconv.Atoi(scanner.Text())
	// scanner.Scan()
	// penduduk1.alamat = scanner.Text()

	// fmt.Printf("Hello, my name is %s. Im %d years old. I live in %s", penduduk1.nama, penduduk1.umur, penduduk1.alamat)
	var murid Student
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	murid.Name = scanner.Text()
	scanner.Scan()
	murid.Age, _ = strconv.Atoi(scanner.Text())

	student1 := Student{Name: murid.Name, Age: murid.Age}
	student1.Introduction()
}

// Quiz Bamboo
// func sisasekat(cut int, element []int) []int {
// 	var result []int
// 	for i := 0; i < cut; i++ {
// 		fmt.Printf("Potongan ke %d\n", i+1)
// 		result := sekat[i] - 1
// 		fmt.Println(result)
// 	}
// 	return result
// }
//element := []int{}
// scanner := bufio.NewScanner(os.Stdin)
// scanner.Scan()
// jmlbambo, _ := strconv.Atoi(scanner.Text())
// // element = banyakbamboo(jmlbambo)
// //fmt.Printf("Banyaknya bamboo %d\n", jmlbambo)
// sekat := []int{}
// for i := 1; i <= jmlbambo; i++ {
// 	scanner.Scan()
// 	input, _ := strconv.Atoi(scanner.Text())
// 	sekat = append(sekat, input)
// }
// // fmt.Println("Isis Sekat", sekat)
// scanner.Scan()
// cut, _ := strconv.Atoi(scanner.Text())
// for k := 0; k < cut; k+2+{
// 	fmt.Printf("Potongan ke %d\n", k+1)
// 	for f := 0; f < len(element); f++ {
// 		element[f] = element[f] - 1
// 		fmt.Println(element[f])
// 	}
// }

// finalresult = append(finalresult, element[k]-1)
// fmt.Println(finalresult)
// Quiz Pointer
// func giveMarble(giver *int, receiver *int, marble int) {
// 	if *giver >= marble {
// 		*giver = *giver - marble
// 		*receiver = *receiver - *giver
// 	} else {
// 		fmt.Println("Kelereng tidak cukup untuk dibagikan")
// 	}
// }
// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	giver, _ := strconv.Atoi(scanner.Text())
// 	scanner.Scan()
// 	receiver, _ := strconv.Atoi(scanner.Text())
// 	scanner.Scan()
// 	marble, _ := strconv.Atoi(scanner.Text())
// 	giveMarble(&giver, &receiver, marble)
// 	fmt.Printf("%d \n", giver)
// 	fmt.Printf("%d", receiver)
// }

// Quiz Map
// month := [12]string{"januari", "februari", "maret", "april", "mei", "juni",
// 		"juli", "agustus", "september", "oktober", "november", "desember"}
// 	penjualan := map[string]int{
// 		"januari":   2836,
// 		"februari":  3282,
// 		"maret":     787,
// 		"april":     6238,
// 		"mei":       1992,
// 		"juni":      824,
// 		"juli":      2903,
// 		"agustus":   602,
// 		"september": 930,
// 		"oktober":   1002,
// 		"november":  922,
// 		"desember":  3219,
// 	}
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	bulan1, _ := strconv.Atoi(scanner.Text())
// 	scanner.Scan()
// 	bulan2, _ := strconv.Atoi(scanner.Text())
// 	var simpanbulan1 string
// 	if bulan1 <= bulan2 && bulan2 <= 12 {
// 		for i := bulan1 - 1; i < bulan2; i++ {
// 			simpanbulan1 = month[i]
// 			fmt.Printf("Bulan %s : %d tusuk\n", simpanbulan1, penjualan[simpanbulan1])
// 		}
// 	} else {
// 		fmt.Println("FALSE BROO")
// 	}

//Quiz Slice
// func evenNames(slice []string) []string {
// 	hasil := []string{}
// 	for i := 0; i < len(slice); i++ {
// 		if len(slice[i])%2 == 0 {
// 			hasil = append(hasil, slice[i])
// 		}
// 	}
// 	return hasil
// }
// func main(){
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	x := scanner.Text()
// 	slice := strings.Split(x, " ")
// 	names := evenNames(slice)
// 	fmt.Println(strings.Join(names, " "))
// }

// Quiz Array
//var kapasitas int
// fmt.Scanln(&kapasitas)
// inputelement := bufio.NewScanner(os.Stdin)
// inputelement.Scan()
// element := inputelement.Text()
// var Array = strings.Split(element, " ")
// for a := 0; a < kapasitas; a++ {
// 	convert, err := strconv.Atoi(Array[a])
// 	if err == nil {
// 		if convert%2 == 0 {
// 			fmt.Println(convert)
// 		}
// 	}
// }
