package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/dvontrec/circuitbreakerserver/pkg/client"
)

func main() {
	fmt.Println("Here I am, running...")
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(net.JoinHostPort("", "9000"), hystrixStreamHandler)
	hystrix.ConfigureCommand("Get_Request", hystrix.CommandConfig{
		Timeout:                500,
		MaxConcurrentRequests:  15,
		ErrorPercentThreshold:  50,
		RequestVolumeThreshold: 10,
		SleepWindow:            30000,
	})
	for {
		_ = hystrix.Do("Get_Request", func() error {
			c := client.New("http://0.0.0:8109/circuit")
			err := c.Request()
			return err
		}, nil)
		time.Sleep(time.Millisecond * 300)
	}
}
