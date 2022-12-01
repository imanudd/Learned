package main

import "fmt"

type Student struct {
	Name   string
	Alamat string
	Nim    int
}

//make a method inheritance of struct Student, siswa is a variable, Student is a struct, and NewStudent is a name of func
func (siswa Student) NewStudent() {
	fmt.Printf("Halo Nama Saya %s Dan Saya Memilki NIM %d\n", siswa.Name, siswa.Nim)
}

func main() {
	// New Variable "murid" for make a new object of struct
	var murid Student
	// New object Name,Alamat,Nim
	murid.Name = "Jamil"
	murid.Alamat = "Jl Albaidho 2"
	murid.Nim = 202243501459
	//call struct Student
	fmt.Println(murid)
	fmt.Printf("Halo Nama Saya %s, Saya tinggal di %s, Saya memiliki Nim %d\n", murid.Name, murid.Alamat, murid.Nim)

	//call func NewStudent
	siswa1 := Student{Name: "Imanuddin", Nim: 202243501459}
	siswa1.NewStudent()
	siswa2 := Student{Name: "Faldini", Nim: 90283992829}
	siswa2.NewStudent()

}
