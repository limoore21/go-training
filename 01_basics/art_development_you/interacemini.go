package main

import "fmt"

type Gadget struct{}

type Player interface {
	Play() string
}

type Volume interface {
	Wave() int
	Music() string
}

type SuperPlayer interface {
	Player
	Volume
}

func (g Gadget) Play() string {
	return "Мяу-мяу, играет кошачья песня"
}

func (g Gadget) Music() string {
	return "Бродяга СОС"
}

func (g Gadget) Wave() int {
	return 1488
}

func main() {
	myGadget := Gadget{}

	var sp SuperPlayer
	sp = myGadget

	fmt.Println(sp.Play())
	fmt.Printf("Wave: %d, Music: %s\n", sp.Wave(), sp.Music())
}
