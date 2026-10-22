package main

import "fmt"

type Stats struct {
	counters map[string]int
}

func (s *Stats) Increment(metric string) {
	s.counters[metric]++
}

func main() {
	var st Stats
	st.Increment("cpu_spike")
	fmt.Println(st.counters["cpu_spike"])
}
