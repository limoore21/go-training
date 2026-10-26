package main

import "fmt"

type MyError struct{}

func (e *MyError) Error() string { return "error" }

func main() {
	var p *MyError = nil
	var err error = p
	fmt.Println(err == nil)
}
