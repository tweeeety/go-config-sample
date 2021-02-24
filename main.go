package main

import "fmt"

func main() {
	path := "./config.yaml"
	c := LoadConfig(path)
	bconf, err := c.GetBqConfig("fuga")
	fmt.Printf("bconf: %+v\n", bconf)
	fmt.Printf("error: %+v\n", err)
}
