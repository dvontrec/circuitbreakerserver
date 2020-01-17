package main

import (
	"fmt"

	"github.com/dvontrec/circuitbreakerserver/pkg/client"
)

func main() {
	fmt.Println("Here I am, running...")
	c := client.New("http://0.0.0:8109/circuit")
	i := 0
	for i < 30 {
		i++
		c.Request()
	}
}
