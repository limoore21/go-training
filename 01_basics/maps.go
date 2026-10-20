package main

import "fmt"

func main() {
	pins := make(map[string]int)

	pins["admin"] = 1234
	pins["guest"] = 0000
	fmt.Println(pins["admin"])
	pins["admin"] = 7777
	fmt.Println(pins["admin"])
}
