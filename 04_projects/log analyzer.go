package main

import "fmt"

type Analyzer struct {
	request map[string]int
	banned  map[string]bool
}

func NewAnalyzer() *Analyzer {
	a := Analyzer{
		request: make(map[string]int),
		banned:  make(map[string]bool),
	}
	return &a
}

func (a *Analyzer) ProcessIP(ip string) {
	a.request[ip]++
	if a.request[ip] > 3 {
		a.banned[ip] = true
	}
}
func (a *Analyzer) PrintReport() {
	for ip, count := range a.request {
		fmt.Printf("IP: %s, Стучался: %d\n", ip, count)
		if a.banned[ip] == true {
			fmt.Println("[Banned]")
		}
	}
}

func main() {
	analyzer := NewAnalyzer()
	logs := []string{
		"192.168.1.20",
		"10.0.0.5",
		"10.0.0.5",
		"10.0.0.5",
		"10.0.0.5",
		"10.0.0.5",
	}
	for _, log := range logs {
		analyzer.ProcessIP(log)
	}
	analyzer.PrintReport()
}
