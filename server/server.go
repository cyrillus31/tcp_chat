package server

import (
	"context"
	// "io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/cyrillus31/tcp_chat/handler"
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

var Handler = handler.Handler{}

func (s Server) processConnection(ctx context.Context, conn net.Conn) {
	for {
		select {
		case <-s.quitCh:
			println("Stopping processConnection")
			return
		default:
			// io.Copy(os.Stdout, conn)
			err := Handler.HandleData(conn, os.Stdout)
			if err != nil {
				panic(err)
			}
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
	ctx := context.Background()
	go s.lookForConnections(ctx)
	signal.Notify(s.quitCh, os.Interrupt, syscall.SIGTERM)
	defer log.Println("Server gracefully stops!")
	select {
	case <-s.quitCh:
		return
	}
}

func (s *Server) Stop() {
	s.listener.Close()
}
