package main

import "fmt"

func doSomething(calback func(int, int) int, s string) int {
	result := calback(5, 1)
	fmt.Println(s)
	return result
}

func main() {
	sumCalback := func(n1, n2 int) int {
		return n1 + n2
	}
	result := doSomething(sumCalback, "plus")
	fmt.Println(result)

	multypleCalback := func(n1, n2 int) int {
		return n1 * n2
	}
	result = doSomething(multypleCalback, "multyple")
	fmt.Println(result)
}
