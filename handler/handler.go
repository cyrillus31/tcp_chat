package handler

import (
	"io"
	"net"
)

const (
	VERSION=byte(100)
	READ_JSON=byte(1)
	HEADER_LENGTH=4
)

type Handler struct {}

// Unmarshallas the message and writes it to the writer
func (h *Handler) HandleData(w *io.Writer, conn *net.Conn) error {
	return nil
}
