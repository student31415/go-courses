package main

import (
	"fmt"
	"time"
)

func main() {
	mails := make(chan string)
	for j := 0; j < 3; j++ {
		go func(i int) {
			for {
				message := <-mails
				SendMail(message, i)
			}
		}(j)
	}
	for i := 0; i < 100; i++ {
		mails <- fmt.Sprintf("mail №-%d", i)
	}
	fmt.Scan()

}

//100 писем
//10 рассыльщиков
//после отправки - ждать 5 сек

func SendMail(message string, i int) {
	mailer := fmt.Sprintf("Mailer #%d send message:%s", i, message)
	fmt.Println(mailer)
	k := time.Duration(30 + i)
	time.Sleep(k * time.Second)
}
