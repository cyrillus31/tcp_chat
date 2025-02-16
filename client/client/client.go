package client

import (
	"net"
	"os"

	"github.com/cyrillus31/tcp_chat/internals/utils"
)

type Client struct {
	connAddr string
	conn     net.Conn
	quitch   chan struct{}
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
		err := utils.Copy(c.conn, os.Stdin)
		if err != nil {
			panic(err)
		}
	}
}

func (c *Client) Start() {
	go c.writeMessage()
	<-c.quitch
}

func (c *Client) Stop() {
	c.conn.Close()
	println("Connection closed")
}
