package main

import "fmt"

type structHere struct {
	N1, N2 int
}

func (s *structHere) sum() int {
	return s.N1 + s.N2
}

type interfaceHere interface {
	sum() int
}

func main() {
	var os *structHere
	var i interfaceHere
	i = os
	//fmt.Println(i.sum())
	fmt.Printf("%v, %T\n", i, i)
	if i == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

type Otherstruct struct {
	A, B int
}

func (o Otherstruct) sum() int {
	return o.A + o.B
}
