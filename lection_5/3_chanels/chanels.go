// Один из механизмов синхронизации - каналы
// Каналы, это объект через который можно обеспечить взаимодействие нескольких горутин
// В принимающей (или возвращающей) канал функции, можно указать направление работы с каналом
// Только для чтения - "<-chan" или только для записи "chan<-"
package main

import (
	"fmt"
)

var c chan int

func main() {
	// Создаем канал
	c := make(chan int)
	// стартуем пишущую горутину
	go greet(c)

	//for i := 0; i < 5; i++ {
	//	// Читаем пару строк из канала
	//	fmt.Println(<-c, ",", <-c)
	//}

	go func(r <-chan int) {
		for i := 0; i < 5; i++ {
			//fmt.Println("start reading")
			j := <-r
			fmt.Println("read from chanel by gorutine1 i :", j)

		}
	}(c)
	go func(r <-chan int) {
		for i := 0; i < 5; i++ {
			//fmt.Println("start reading")
			j := <-r
			fmt.Println("read from chanel by gorutine2 i :", j)

		}
	}(c)
	fmt.Scanln()
	//stuff := make(chan int, 7)
	//for i := 0; i < 19; i = i + 3 {
	//	stuff <- i
	//}
	//close(stuff)
	//fmt.Println("Res", process(stuff))
}

func greet(c chan<- int) {
	var i int
	// Запускаем бесконечный цикл
	for {
		// и пишем в канал пару строк
		// Подпрограмма будет заблокирована до того, как кто-то захочет прочитать из канала
		//c <- fmt.Sprintf("Владыка")
		//c <- fmt.Sprintf("Штурмовик")
		//i++
		//fmt.Println("iteration:",i)
		i++
		//	fmt.Println("start writing")
		c <- i
		fmt.Println("add to chanel i :", i)
	}
}

func process(input <-chan int) (res int) {
	for r := range input {
		fmt.Printf("Res:%d ; r:%d\n", res, r)
		res += r

	}
	return
}
