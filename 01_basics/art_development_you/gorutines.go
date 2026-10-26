package main

import (
	"fmt"
	"time"
)

func say(word string) {
	fmt.Println(word)
}

func sayHello(exit chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		exit <- i
	}
	close(exit)
}

func main() {
	ch := make(chan int)
	go sayHello(ch)
	for i := range ch {
		fmt.Println(i)
	}
}
