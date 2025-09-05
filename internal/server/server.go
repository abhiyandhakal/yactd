package server

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
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

	// Read the incoming data
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		res := protocol.Response{
			Status: "error",
			Error:  fmt.Sprintf("error reading request: %v", err),
		}
		json.NewEncoder(conn).Encode(res)
		return
	}

	// Parse the command and arguments
	cmd := string(buf[:n])
	parts := strings.Fields(cmd)
	if len(parts) < 1 {
		res := protocol.Response{
			Status: "error",
			Error:  "no command provided",
		}
		json.NewEncoder(conn).Encode(res)
		return
	}

	action := parts[0]
	payload := make(map[string]string)

	// For pull command, expect format: pull <image>
	if action == "pull" && len(parts) > 1 {
		payload["image"] = parts[1]
	}

	// Create and dispatch the request
	req := protocol.Request{
		Action:  action,
		Payload: payload,
	}

	res := handler.Dispatch(req)
	json.NewEncoder(conn).Encode(res)
}
