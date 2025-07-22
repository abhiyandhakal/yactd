package server

import (
	"encoding/json"
	"net"
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
	defer listener.Close()
	for {
		conn, err := listener.Accept()
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
