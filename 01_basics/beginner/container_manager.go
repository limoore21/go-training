package main

import "fmt"

type Container struct {
	ID    string
	State string
	Image string
}

func NewContainer(id, image string) *Container {
	newContainer := Container{
		ID:    id,
		State: "created",
		Image: image,
	}
	return &newContainer
}
func (c *Container) Start() {
	c.State = "running"
}

func (c *Container) Stop() {
	c.State = "exited"
}

func main() {
	c := NewContainer("777", "nginx:latest")
	fmt.Printf("Итоговый статус: %s, Образ: %s, ID: %s\n", c.State, c.Image, c.ID)
	c.Start()
	fmt.Printf("Итоговый статус: %s, Образ: %s, ID: %s\n", c.State, c.Image, c.ID)
	c.Stop()
	fmt.Printf("Итоговый статус: %s, Образ: %s, ID: %s\n", c.State, c.Image, c.ID)
}
