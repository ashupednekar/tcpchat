package server

import (
	"fmt"
	"log"
	"net"

	"github.com/ashupednekar/tcpchat/chat"
	"gorm.io/gorm"
)

// TODO: figure out a way to use composition with ashupednekar/gotcp package

type Channels struct {
	UserInputMap   map[string](chan string)
	GrpReceiveChan map[string](chan string)
	Msgchan        chan Message
	quitchan       chan struct{}
}

type Server struct {
	ListenAddr string
	ln         net.Listener
	Chans      Channels
	Db         *gorm.DB
}

type Message struct {
	Source  string
	Payload []byte
}

func NewServer(addr string) *Server {
	return &Server{
		ListenAddr: addr,
		Chans: Channels{
			UserInputMap:   make(map[string](chan string)),
			GrpReceiveChan: make(map[string](chan string)),
			quitchan:       make(chan struct{}),
		},
		Db: chat.GetDb(),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.ln = ln
	s.AcceptLoop()

	<-s.Chans.quitchan
	for _, c := range s.Chans.UserInputMap {
		close(c)
	}
	for g, c := range s.Chans.GrpReceiveChan {
		c <- fmt.Sprintf("Hey %s Server is closing, goodbye for now", g)
		close(c)
	}
	return nil
}

func (s *Server) AcceptLoop() {
	println("Accepting connections at ", s.ListenAddr)
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Fatal("error while accepting: ", err)
		}
		go s.HandleConn(conn)
	}
}
