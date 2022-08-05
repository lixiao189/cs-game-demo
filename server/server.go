package server

// TODO Closing connection gracefully

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/lixiao189/cs-game-demo/protocol"
	"github.com/lixiao189/cs-game-demo/util"
	"github.com/tidwall/gjson"
	"github.com/xtaci/kcp-go/v5"
)

type Server struct {
	Host string
	Port int

	Connections map[string]net.Conn
}

func (s *Server) ServerInit() {
	log.Println("Running on server mode")

	laddr := fmt.Sprintf("%v:%v", s.Host, s.Port)
	listener, err := kcp.Listen(laddr)
	if err != nil {
		defer listener.Close()
		log.Fatal(err)
		os.Exit(-1)
	}

	// Listenning loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go s.handlePacket(conn)
	}
}

func (s *Server) handlePacket(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		util.HandleErr(err)
		packet := string(buf)
		log.Println(packet)

		switch gjson.Get(packet, "type").String() {
		case protocol.PlayerJoinType:
			newPlayerName := gjson.Get(packet, "data.name").String()
			log.Println(newPlayerName + " joins the game")

			// Add connection to server
			s.Connections[newPlayerName] = conn

			// TODO broadcast new player's space ship info
			
		}
	}
}
