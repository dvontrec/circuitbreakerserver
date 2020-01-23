package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type response struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

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
func (c *Client) Request() error {
	resp, err := http.Get(c.url)
	if err != nil {
		return err
	}
	var res response
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return err
	}
	fmt.Printf("*****************\n%+v\n", res)
	if resp.StatusCode != 200 {
		return errors.New("Didnt work")
	}
	defer resp.Body.Close()
	return nil
}
