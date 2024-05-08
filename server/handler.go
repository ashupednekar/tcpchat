package server

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/ashupednekar/tcpchat/chat/data_adapters/mutators"
	"github.com/ashupednekar/tcpchat/chat/data_adapters/selectors"
)

func (s *Server) GetChan(IP string) chan string {
	if ch, ok := s.ChanMap[IP]; ok {
		return ch // Channel exists, return it
	}

	// Channel doesn't exist, create a new one
	newCh := make(chan string)
	s.ChanMap[IP] = newCh
	return newCh
}

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

		ch := s.GetChan(conn.RemoteAddr().String())

		go func() {
			for s := range ch {
				fmt.Println("to send: ", s)
				fmt.Fprintf(conn, s)
			}
		}()

		switch {
		case strings.HasPrefix(string(msg), "root:join"):
			HandleJoin(*s, string(msg), conn)
		case strings.HasPrefix(string(msg), "chat:"):
			HandleChat(*s, string(msg), conn)
		default:
			fmt.Fprintf(conn, "invalid payload received, skipping")
		}

	}
}

func HandleJoin(s Server, msg string, conn net.Conn) {
	l := strings.Split(string(msg), ":")
	Name := l[len(l)-2]
	fmt.Println("New user joining: ", Name)
	err := mutators.CreateUser(s.Db, Name, conn.RemoteAddr().String())
	if err != nil {
		fmt.Fprintf(conn, "error creating user: ", err)
	}
}

func HandleChat(s Server, msg string, conn net.Conn) {
	err, user := selectors.GetUser(s.Db, conn.RemoteAddr().String())
	if err != nil {
		fmt.Fprintf(conn, "error retrieving user: ", err)
	}
	l := strings.Split(string(msg), ":")
	receiverName := l[1]
	text := l[len(l)-2]
	fmt.Printf("received message from user: %s => %s\n", user.Name, string(msg))
	mutators.SaveTextMessage(s.Db, conn.RemoteAddr().String(), receiverName, text)
	err, IPs := selectors.GetIPsFromGroupName(s.Db, receiverName)
	if err != nil {
		fmt.Fprintf(conn, "group or individual %s not found", receiverName)
	}

	for _, ip := range IPs {
		recv := s.GetChan(ip)
		recv <- text
	}
}
