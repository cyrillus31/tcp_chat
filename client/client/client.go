package client

import (
	"fmt"
	"net"
)

type Client struct {
	connAddr string
	conn net.Conn
}

func New(connAddr string) *Client {
	conn, err := net.Dial("tcp", connAddr)
	if err != nil {
		panic(err)
	}
	return &Client{
		connAddr: connAddr,
		conn: conn,
	}
}

func (c *Client) Start() {
	for {
		input := []byte{}
		fmt.Scan(&input)
		_, err := c.conn.Write(input)
		if err != nil {
			return
		}
		println("Message sent")
	}
}

func (c *Client) Stop() {
	c.conn.Close()
	println("Connection closed")
}
