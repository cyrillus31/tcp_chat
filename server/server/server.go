package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
)

type Server struct {
	listnAddr string
	listener  net.Listener
}

func NewServer(listAddr string) *Server {
	return &Server{
		listnAddr: listAddr,
	}
}

func (s Server) processConnection(ctx context.Context, conn net.Conn) {
	for {
		select {
		case <-ctx.Done():
			println("Stopping processConnection")
			return
		default:
		io.Copy(os.Stdout, conn)
		}
	}
}
func (s Server) lookForConnections(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			println("Stopping lookForConnections")
			return
		default:
			println("Looking for connection")
			conn, err := s.listener.Accept()
			if err != nil {
				panic(err)
			}
			println("Connection created to", conn.RemoteAddr().String())
			go s.processConnection(ctx, conn)
			defer conn.Close()
		}
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.listnAddr)
	if err != nil {
		panic(err)
	}
	s.listener = listener
	ctx, cancel := context.WithCancel(context.Background())
	go s.lookForConnections(ctx)
	for {
		var input string
		fmt.Scan(&input)
		if input == "stop" {
			cancel()
			fmt.Println("Connections cancelled!")
		}
	}
}

func (s *Server) Stop() {
	s.listener.Close()
}
