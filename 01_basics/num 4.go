package main

import "fmt"

func main() {
	port := 80
	if true {
		port = 443
	}
	fmt.Println(port)
}
