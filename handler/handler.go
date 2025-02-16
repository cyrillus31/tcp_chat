package handler

import (
	"fmt"
	"io"
	"net"
	"time"
)

const (
	VERSION       = byte(100)
	READ_JSON     = byte(1)
	HEADER_LENGTH = 4
)

type Handler struct{}

// Read Message from r and write it to conn
func (h *Handler) SendData(conn net.Conn, r io.Reader) error {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}
	msg := Message{User: conn.LocalAddr().String(), Text: string(buf[:n])}
	msgData, err := msg.Marshall()
	if err != nil {
		return err
	}
	conn.Write(msgData)
	return nil
}


// Get Message from conn and write it to w
func (h *Handler) HandleData(conn net.Conn, w io.Writer) error {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return err
	}
	msg := Message{}
	err = msg.Unmarshall(buf[:n])
	if err != nil {
		return err
	}
	msgText := fmt.Sprintf("From: %s\tAt: %s\nText: %s\n", msg.User, time.Now().Format("15:04:05"), msg.Text)
	w.Write([]byte(msgText))
	return nil
}
