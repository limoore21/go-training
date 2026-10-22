package main

import "fmt"

func pointers() {
	a := "hello world"
	b := 42
	fmt.Println(a)
	fmt.Println(b)
	p := &a
	// p содержит ссылку на ячейку переменной а
	fmt.Println(*p)
	// * позволяет заглянуть/редактировать ячейку
}

func main() {
	pointers()
}
