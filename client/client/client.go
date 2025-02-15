package client

import (
	"io"
	"net"
	"os"
	"sync"
)

type Client struct {
	connAddr string
	conn     net.Conn
}

func New(connAddr string) *Client {
	conn, err := net.Dial("tcp", connAddr)
	if err != nil {
		panic(err)
	}
	return &Client{
		connAddr: connAddr,
		conn:     conn,
	}
}

func (c *Client) writeMessage() {
	for {
		// var input []byte
		// fmt.Scan(&input)
		// _, err := c.conn.Write(input)
		// if err != nil {
		// 	return
		// }
		io.Copy(c.conn, os.Stdin)
	}
}

func (c *Client) Start() {
	var wg sync.WaitGroup
	wg.Add(1)
	go c.writeMessage()
	wg.Wait()
}

func (c *Client) Stop() {
	c.conn.Close()
	println("Connection closed")
}
