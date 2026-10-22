package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) movePoint(x, y int) Point {
	p.X += x
	p.Y += y
	return *p
}

func (p *Point) movePointPtr(x, y int) {
	p.X += x
	p.Y += y
}

func main() {
	p := Point{1, 2}
	//fmt.Println(p)
	//fmt.Println(p.movePoint(p, 1, 1))
	//fmt.Println(p)
	p.movePointPtr(1, 1)
	//fmt.Println(p)
	v := &p
	v.movePoint(1, 2)
	fmt.Println(p)
	v.movePointPtr(1, 3)
	fmt.Println(p)

}
