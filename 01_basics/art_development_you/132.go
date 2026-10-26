package main

import "fmt"

type Person struct {
	name string
	pass string
}

func (p *Person) String() string {
	return fmt.Sprintf("{name: %s, pass: [REDACTED]}", p.name)
}

func main() {
	p := &Person{
		name: "Bob",
		pass: "123456",
	}
	fmt.Println(p)
}
