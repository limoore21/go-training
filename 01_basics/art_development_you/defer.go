package main

import "fmt"

func main() {
	// stack LFO Last In First Out
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("hello")
}
