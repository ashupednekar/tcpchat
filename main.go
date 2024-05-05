package main

import "github.com/ashupednekar/gotcp/server"

func main() {
	server := server.NewServer(":3000")
	server.Start()
}
