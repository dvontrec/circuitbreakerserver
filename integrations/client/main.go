package main

import (
	"fmt"
	"time"

	"github.com/dvontrec/circuitbreaker/pkg/resiliency"
	"github.com/dvontrec/circuitbreakerserver/pkg/client"
)

type queue []int

func main() {
	myCircuit := resiliency.CircuitBreaker{
		Name:             "Get Request",
		FailureThreshold: 5,
	}
	fmt.Println("Here I am, running...")
	myQue := queue{}
	i := 0
	for {
		fmt.Println("Cir State: ", myCircuit.GetStatus())
		url := fmt.Sprintf("http://0.0.0:8109/circuit?value=%v", i)
		myCircuit.RunFunction(func() error {
			if len(myQue) != 0 {
				i, myQue = myQue.deque()
			}
			c := client.New(url)
			err := c.Request()
			if err == nil {
				i++
			}
			return err

		})
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
