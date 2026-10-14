package main

import "fmt"

type Config struct {
	Port      int
	DebugMode bool
}

func main() {
	config := Config{
		Port:      80,
		DebugMode: false,
	}
	enableDebug(&config)
	fmt.Printf("config: %v\n", config)

}
func enableDebug(c *Config) {
	c.DebugMode = true
	c.Port = 8080
}
