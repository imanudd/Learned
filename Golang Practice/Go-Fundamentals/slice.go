package main

import "fmt"

func main() {
	// array := [4]string{"Utara", "Timur", "Barat", "Selatan"}
	// var slice = array[0:2] // slice array from 0 to 2

	// fmt.Printf("%v Data Type %T\n", array, array)
	// fmt.Printf("%v Data Type %T\n", slice, slice)

	array := []string{"Utara", "Timur", "Barat", "Selatan"}
	fmt.Println(array)
	slicearray := array[2:] // slice array form 2 to end
	//slice
	fmt.Printf("%v Data Type %T\n", array, array)
	fmt.Printf("%v Data Type %T\n", slicearray, slicearray)

	//fungsi len(mencari jumlah isi dari slice/array) dan cap(mencari total kapasitas dari slice/array) di slice
	fmt.Printf("isi slice %v , jumlah element slice = %d, kapasitas slice %d\n", array, len(array), cap(array))
	fmt.Printf("isi slice %v , jumlah element slice = %d, kapasitas slice %d\n", slicearray, len(slicearray), cap(slicearray))

	// to add element in array with APPEND
	newinputarray := append(slicearray, "Barat Laut", "timur Laut", "Tenggara")
	fmt.Printf("isi slice %v , jumlah element slice = %d, kapasitas slice %d \n", newinputarray, len(newinputarray), cap(newinputarray))

	// code to make array with no value or null element, and you can add with slice or append
	arraymake := make([]int, 3)
	fmt.Println("array", arraymake)

	// untuk mencari lokasi penyimpanan array dan slice dan untuk membandingkan array dan slice
	fmt.Printf(" array Located At = %p\n", &array[0])
	fmt.Printf(" slice array Located At = %p\n", &slicearray[0])
	fmt.Printf(" newinputarray Located At = %p\n", &newinputarray[0])

	//Copy Slice Untuk Ke Slice Lain
	dst := make([]string, 3)
	src := []string{"Semangka", "Mangga", "Lemon"}
	copy(dst, src)

	fmt.Printf("Isi SRC Slice %v , located At %p\n", src, src)
	fmt.Printf("Isi dst Slice %v , located At %p\n", dst, dst)

}
