package server

import (
	"net"
	"os"
)

type Server struct {
	listnAddr string
	listener net.Listener
}

func NewServer(listAddr string) *Server {
	return &Server{
		listnAddr: listAddr,
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.listnAddr)
	if err != nil {
		panic(err)
	}
	s.listener = listener
	conn, err := s.listener.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for {
		buf := make([]byte, 8)
		conn.Read(buf)
		os.Stdout.Write(buf)
	}
}

func (s *Server) Stop() {
	s.listener.Close()
}
