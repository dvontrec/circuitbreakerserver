package main

import (
	"fmt"

	"github.com/dvontrec/circuitbreakerserver/internal/restapi"
)

func main() {
	api := restapi.New()
	fmt.Println("Server Listening on port 8109")
	api.Start()
}
