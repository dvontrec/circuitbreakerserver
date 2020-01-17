package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/dvontrec/circuitbreakerserver/pkg/client"
)

func main() {
	fmt.Println("Here I am, running...")
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "81"), hystrixStreamHandler)
	hystrix.ConfigureCommand("Get_Request", hystrix.CommandConfig{
		Timeout:                500,
		MaxConcurrentRequests:  5,
		ErrorPercentThreshold:  25,
		RequestVolumeThreshold: 3,
		SleepWindow:            30000,
	})
	for {
		_ = hystrix.Do("Get_Request", func() error {
			c := client.New("http://0.0.0:8109/circuit")
			err := c.Request()
			return err
		}, nil)
	}
}
