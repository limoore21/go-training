package main

import (
	"fmt"
	"time"
)

func sendEmail(email string, ch chan string) {
	time.Sleep(1500 * time.Millisecond) // Долгая отправка email
	ch <- fmt.Sprintf("Email успешно отправлен на %s", email)
}

func sendSMS(phone string, ch chan string) {
	time.Sleep(1000 * time.Millisecond) // Долгая отправка SMS
	ch <- fmt.Sprintf("SMS успешно отправлено на %s", phone)
}

func main() {
	startTime := time.Now()
	ch := make(chan string)

	// Запускаем отправку параллельно
	go sendEmail("bro@example.com", ch)
	go sendSMS("+79991112233", ch)

	emailResult := <-ch
	smsResult := <-ch

	for i := 0; i < 2; i++ {
		fmt.Println(emailResult)
		fmt.Println(smsResult)
	}

	fmt.Println("Затрачено времени:", time.Since(startTime))
}
