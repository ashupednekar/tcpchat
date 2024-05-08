package main

import (
	"github.com/ashupednekar/gotcp/server"
	"github.com/ashupednekar/tcpchat/chat"
)

func main() {
	db := chat.GetDb()
	chat.Migrate(db)
	server := server.NewServer(":3000")
	server.Start()
}
