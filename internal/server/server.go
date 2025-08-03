package server

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"yactd/internal/handler"
	"yactd/internal/protocol"
)

type Server struct {
	socketPath string
}

func New(socketPath string) *Server {
	return &Server{socketPath: socketPath}
}

func (s *Server) Start() {
	listener, err := net.Listen("unix", s.socketPath)

	if err != nil {
		panic(err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
		listener.Close()
	}()
	for {
		conn, err := listener.Accept()
		select {
		case <-done:
			fmt.Println("Yact Daemon sutting down...")
			return
		default:
		}
		if err != nil {
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()
	var req protocol.Request
	res := handler.Dispatch(req)
	json.NewEncoder(conn).Encode(res)
}
