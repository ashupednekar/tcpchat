package main

import (
	"github.com/ashupednekar/tcpchat/chat"
)

func main() {
	// server := server.NewServer(":3000")
	// server.Start()
	db := chat.GetDb()
	chat.Migrate(db)
}
