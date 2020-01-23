package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/dvontrec/circuitbreakerserver/pkg/client"
)

type queue []int

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
	myQue := queue{}
	i := 0
	for {
		fmt.Println("Queue Len: ", len(myQue))
		url := fmt.Sprintf("http://0.0.0:8109/circuit?value=%v", i)
		_ = hystrix.Do("Get_Request", func() error {
			if len(myQue) != 0 {
				i, myQue = myQue.deque()
			}
			c := client.New(url)
			err := c.Request()
			if err != nil {
				myQue.que(i)
			}
			i++
			return err
		}, nil)
		time.Sleep(time.Millisecond * 300)
	}
}

func (q *queue) que(val int) {
	*q = append(*q, val)
}

func (q queue) deque() (int, queue) {
	x, q := q[len(q)-1], q[:len(q)-1]
	return x, q
}
