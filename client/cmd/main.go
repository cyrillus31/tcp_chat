package main

import "github.com/cyrillus31/tcp_chat/client"

func main() {
	client := client.New("0.0.0.0:3000")
	client.Start()
	defer client.Stop()
}
