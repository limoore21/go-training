package main

import "fmt"

type Route struct {
	IP      string
	Blocked bool
}

func main() {
	ips := []Route{
		{IP: "127.0.0.1", Blocked: true},
		{IP: "127.0.0.2", Blocked: true},
		{IP: "127.0.0.3", Blocked: true},
	}
	for index, _ := range ips {
		applyRules(&ips[index])
		fmt.Printf("IP: %s, Статус: %t\n", ips[index].IP, ips[index].Blocked)
	}
}

func applyRules(r *Route) {
	if r.IP == "1.1.1.1" {
		r.Blocked = false
	} else {
		r.Blocked = true
	}
}
