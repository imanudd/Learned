package main

import "fmt"

type Siswa interface {
	StudentName() string
}

func GetStudent(murid Siswa) {
	fmt.Println("helo", Siswa.StudentName)
}

type Biodata struct {
	Name string
}

func (biodata Biodata) StudentName() string {
	return biodata.Name
}

func main() {
	var Siswa1 Biodata
	Siswa1.Name = "Jamil"
	GetStudent(&Siswa1)
}
