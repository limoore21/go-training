package main

import "fmt"

type Backend struct {
	Host   string
	Load   int
	Status string
}

func analyzeLoad(b *Backend) {
	if b.Load > 95 {
		b.Status = "Critical"
	} else if b.Load > 80 {
		b.Status = "WARNING"
	}
}

func main() {
	cluster := []Backend{
		{Host: "web-1", Load: 45, Status: "Normal"},
		{Host: "web-2", Load: 85, Status: "Normal"},
		{Host: "web-3", Load: 99, Status: "Normal"},
	}
	for index, _ := range cluster {
		analyzeLoad(&cluster[index])
	}
	fmt.Println(cluster)
}
