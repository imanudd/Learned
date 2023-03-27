package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type BankAcc struct {
	Balance int
	Mutex   sync.Mutex
	Wege    sync.WaitGroup
}

func main() {
	accountBalance := BankAcc{Balance: 0}

	accountBalance.Wege.Add(10)
	r := rand.New(rand.NewSource(99))
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			go accountBalance.TopUp(r.Intn(100))
		} else {
			go accountBalance.Credit(r.Intn(100))
		}
	}
	accountBalance.Wege.Wait()
	fmt.Println("Total Balance : ", accountBalance.Balance)
}

func (acc *BankAcc) TopUp(amount int) {
	acc.Mutex.Lock()

	acc.Balance += amount

	fmt.Println("aksi topup", amount)
	fmt.Println("Balance = ", acc.Balance)

	acc.Mutex.Unlock()
	acc.Wege.Done()
}

func (acc *BankAcc) Credit(amount int) {
	acc.Mutex.Lock()
	var kalkulasi = acc.Balance - amount
	if kalkulasi >= 0 {
		acc.Balance -= amount
		fmt.Println("aksi credit", amount)
		fmt.Println("Balance = ", acc.Balance)
	} else {
		fmt.Println("aksi credit", amount)
		fmt.Println("Balance = ", acc.Balance)
	}
	acc.Mutex.Unlock()

	acc.Wege.Done()

}
