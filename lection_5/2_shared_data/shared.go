// Go рутины, одновременно работающие с общими данными сами собой не могу синхронизироваться
package main

import (
	"fmt"
	"log"
)

// Пусть у нас есть Счет
type Account struct {
	balance float64
}

func (a *Account) Balance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) {
	log.Printf("depositing: %f", amount)
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) {
	if amount > a.balance {
		log.Printf("withdrawing: balance 0")
		//go a.Withdraw(amount)
		return
	}
	log.Printf("withdrawing: %f", amount)
	a.balance -= amount
}

func main() {
	//acc := Account{}
	//fmt.Println(acc.Balance())
	//// Стартуем 10 go рутин
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		// Каждая из которых, производит операции с аккаунтом
	//		for j := 0; j < 10; j++ {
	//			// Иногда снимает деньги
	//			if j%2 == 1 {
	//				acc.Withdraw(50)
	//				continue
	//			}
	//			// иногда кладет
	//			acc.Deposit(50)
	//		}
	//	}()
	//}
	////fmt.Println("Balance:")
	//
	////time.Sleep(500*time.Microsecond)
	//fmt.Scanln()
	//// Чтоже получится в результате
	//fmt.Println(acc.Balance())

	//var i int
	//fmt.Println(i)
	//
	//for j:=0;j<10000;j++{
	//	go func(){
	//		i=i+1
	//	}()
	//}
	//fmt.Scanln()
	//fmt.Print("result:")
	//fmt.Println(i)
	closure()
	fmt.Scanln()
}

// //
func closure() {
	// Функции захватывают переменные в области видимости
	// Но, чтобы передавать значение, требуется явно передавать его в функцию
	for i := 0; i < 10; i++ {
		go func(a int) {
			fmt.Println("Got", a)
		}(i)
	}
	fmt.Scanln()
}
