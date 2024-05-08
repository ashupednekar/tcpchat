package server

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/ashupednekar/tcpchat/chat/data_adapters/mutators"
	"github.com/ashupednekar/tcpchat/chat/data_adapters/selectors"
	"gorm.io/gorm"
)

func (s *Server) HandleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal("error handling conn: ", err)
			continue
		}
		msg := buf[:n]
		println("msg: ", string(msg))

		switch {
		case strings.HasPrefix(string(msg), "root:join"):
			HandleJoin(s.Db, string(msg), conn)
		case strings.HasPrefix(string(msg), "chat:"):
			HandleChat(s.Db, string(msg), conn)
		default:
			fmt.Fprintf(conn, "invalid payload received, skipping")
		}
	}
}

func HandleJoin(db *gorm.DB, msg string, conn net.Conn) {
	l := strings.Split(string(msg), ":")
	Name := l[len(l)-2]
	fmt.Printf("New user joining: |%s| ", Name)
	err := mutators.CreateUser(db, Name, conn.RemoteAddr().String())
	if err != nil {
		fmt.Fprintf(conn, "error creating user: ", err)
	}
}

func HandleChat(db *gorm.DB, msg string, conn net.Conn) {
	err, user := selectors.GetUser(db, conn.RemoteAddr().String())
	if err != nil {
		fmt.Fprintf(conn, "error retrieving user: ", err)
	}
	fmt.Printf("received message from user: %s => %s\n", user.Name, string(msg))
	l := strings.Split(string(msg), ":")
	mutators.SaveTextMessage(db, conn.RemoteAddr().String(), l[1], l[len(l)-2])
}
