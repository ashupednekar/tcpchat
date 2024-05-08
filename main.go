package main

import (
	"github.com/ashupednekar/tcpchat/chat"
	"github.com/ashupednekar/tcpchat/server"
)

func main() {
	db := chat.GetDb()
	chat.Migrate(db)
	server := server.NewServer(":3000")
	server.Start()
}
