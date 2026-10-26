package main

import "fmt"

func main() {
	var a interface{} = 3.14

	switch a.(type) {
	case int:
		fmt.Println("a is int")
	case string:
		fmt.Println("a is string")
	case bool:
		fmt.Println("a is bool")
	default:
		fmt.Printf("unknown type, type = %T\n", a)
	}
}
