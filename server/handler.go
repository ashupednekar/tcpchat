package server

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/ashupednekar/tcpchat/chat/data_adapters/mutators"
	"github.com/ashupednekar/tcpchat/chat/data_adapters/selectors"
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
			l := strings.Split(string(msg), ":")
			Name := l[len(l)-1]
			fmt.Println("New user joining: ", Name)
			err := mutators.CreateUser(s.Db, Name, conn.RemoteAddr().String())
			if err != nil {
				fmt.Fprintf(conn, "error creating user: ", err)
			}
		case strings.HasPrefix(string(msg), "chat:"):
			err, user := selectors.GetUser(s.Db, conn.RemoteAddr().String())
			if err != nil {
				fmt.Fprintf(conn, "error retrieving user: ", err)
			}
			fmt.Printf("received message from user: %s => %s\n", user.Name, string(msg))
			l := strings.Split(string(msg), ":")
			mutators.SaveTextMessage(s.Db, conn.RemoteAddr().String(), l[1], l[len(l)-1])
		default:
			fmt.Fprintf(conn, "invalid payload received, skipping")
		}
	}
}
