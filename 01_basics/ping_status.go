package main

import "fmt"

func main() {
	ips := []string{"127.0.0.1", "192.168.0.1", "10.0.0.5"}
	for _, ip := range ips {
		currentStatus := "Pending"
		checkPing(ip, &currentStatus)
		fmt.Printf("Адрес: %s %-10s | Статус: %s\n", ip, currentStatus)
	}
}

func checkPing(ip string, status *string) {
	if ip == "192.168.0.1" {
		*status = "Timeout"
	} else {
		*status = "OK"
	}
}
