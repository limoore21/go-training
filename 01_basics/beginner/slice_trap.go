package main

import "fmt"

func UpdatePorts(ports []int) {
	ports[0] = 8080
	ports = append(ports, 9000)
}

func main() {
	serverPorts := []int{80, 443}
	UpdatePorts(serverPorts)
	fmt.Println(serverPorts)
}
