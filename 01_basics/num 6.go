package main

import "fmt"

type Client struct {
	IP       string
	ReqCount int
	IsBanned bool
}

func processRequest(c *Client) {
	c.ReqCount++
	if c.ReqCount == 5 {
		c.IsBanned = true
	}
}

func main() {
	cli := []Client{
		{IP: "192.168.0.1", ReqCount: 0, IsBanned: false},
		{IP: "192.168.0.2", ReqCount: 4, IsBanned: false},
		{IP: "192.168.0.3", ReqCount: 2, IsBanned: false},
	}

	for i := 0; i < 3; i++ {
		for index := range cli {
			processRequest(&cli[index])
		}
	}

	for index := range cli {
		fmt.Printf("IP: %-15s | Запросов: %d | Бан: %v\n",
			cli[index].IP, cli[index].ReqCount, cli[index].IsBanned)
	}
}
