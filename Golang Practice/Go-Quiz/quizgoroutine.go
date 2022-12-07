package main

import (
	"fmt"
	"sync"
)

type DataTest struct {
	User  *User
	Order *Order
}

type User struct {
	Username string
	Email    string
}

type Order struct {
	OrderId    string
	TotalOrder int64
}

var wge sync.WaitGroup

func main() {

	var datatest DataTest
	wge.Add(2)
	go IsiUser(&datatest)
	go IsiOrder(&datatest)
	wge.Wait()
	fmt.Println(*datatest.User)
	fmt.Println(*datatest.Order)

}

func IsiUser(datatest *DataTest) {
	defer wge.Done()
	datatest.User = &User{
		Username: "Jamil",
		Email:    "Iman123@gmail.com",
	}

}

func IsiOrder(datatest *DataTest) {
	defer wge.Done()
	datatest.Order = &Order{
		OrderId:    "001",
		TotalOrder: 10000,
	}
}
