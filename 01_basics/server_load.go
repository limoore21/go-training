package main

import (
	"fmt"
)

type Server struct {
	Host     string
	CpuLoad  int
	isOnline bool
}

func checkOverload(s Server) {
	if s.CpuLoad > 80 {
		fmt.Printf("Сервер: %d\nперегружен!!!", s.CpuLoad)
	} else {
		fmt.Printf("Сервер: %d\n Работает нормально", s.CpuLoad)
	}
}

func main() {
	Myserver := Server{
		Host:    "production-01",
		CpuLoad: 85,
	}
	checkOverload(Myserver)
}
