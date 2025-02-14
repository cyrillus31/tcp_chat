package main

import "github.com/cyrillus31/tcp_chat/server"

func main() {
	server := server.NewServer("0.0.0.0:3000")
	println("Server started!")
	server.Start()
	defer server.Stop()
	println("Server stopped!")
}
