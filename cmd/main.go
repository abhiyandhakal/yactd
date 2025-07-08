package main

import (
	"yactd/internal/server"
)

func main() {
	srv := server.New("/tmp/yactd.sock")
	srv.Start()
}
