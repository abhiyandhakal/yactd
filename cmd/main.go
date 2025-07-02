package main

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(conn *net.UnixConn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Printf("Received: %s\n", string(buf[:n]))

	_, err = conn.Write([]byte("Hello from yactd!\n"))
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
}

func main() {
	socketPath := "/tmp/yactd.sock"
	os.Remove(socketPath) // Remove old socket file if exists

	addr, err := net.ResolveUnixAddr("unix", socketPath)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	listener, err := net.ListenUnix("unix", addr)
	if err != nil {
		fmt.Println("Error listening on Unix socket:", err)
		return
	}
	defer listener.Close()
	defer os.Remove(socketPath)

	fmt.Println("YACTD is listening on", socketPath)

	for {
		conn, err := listener.AcceptUnix()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Accepted connection")

		go handleConnection(conn)
	}
}
