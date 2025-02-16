package server

import (
	"context"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	listnAddr string
	listener  net.Listener
	quitCh    chan os.Signal
}

func NewServer(listAddr string) *Server {
	return &Server{
		listnAddr: listAddr,
		quitCh:    make(chan os.Signal),
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
func (s Server) lookForConnections(ctx context.Context) error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}
		println("Connection created to", conn.RemoteAddr().String())
		go s.processConnection(ctx, conn)
		defer conn.Close()
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.listnAddr)
	if err != nil {
		panic(err)
	}
	s.listener = listener
	defer s.listener.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go s.lookForConnections(ctx)
	signal.Notify(s.quitCh, os.Interrupt, syscall.SIGTERM)
	defer log.Println("Server gracefully stops!")
	select {
	case <-s.quitCh:
		return
	case <-ctx.Done():
		return
	}
}

func (s *Server) Stop() {
	s.listener.Close()
}
