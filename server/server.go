package server

// TODO Closing connection gracefully

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/lixiao189/cs-game-demo/util"
	"github.com/xtaci/kcp-go/v5"
)

type Server struct {
	Host string
	Port int

	Listener net.Listener
	Connections map[string]net.Conn
	WG sync.WaitGroup
	KeyPressChan chan string
}

func (s *Server) ServerInit() {
	log.Println("Running on server mode")

	laddr := fmt.Sprintf("%v:%v", s.Host, s.Port)
	listener, err := kcp.Listen(laddr)
	util.HandleErr(err)
	s.Listener = listener

	s.KeyPressChan = make(chan string, 1000)

	s.initDaemon()
}