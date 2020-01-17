package client

import (
	"fmt"
	"net/http"
)

// Client gives access to client request functionality
type Client struct {
	url string
}

// New returns a new instance of the Client
func New(url string) *Client {
	return &Client{
		url,
	}
}

// Request sends a GET request to recieve data from the Client's url
func (c *Client) Request() {
	resp, err := http.Get(c.url)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}
