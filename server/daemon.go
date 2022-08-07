package server

import (
	"log"
	"net"
	"time"

	"github.com/lixiao189/cs-game-demo/protocol"
	"github.com/lixiao189/cs-game-demo/util"
	"github.com/tidwall/gjson"
)

func (s *Server) initDaemon() {
	s.WG.Add(3)
	{
		go s.listenDaemon()
		go s.broadcastSpaceships()
		go s.broadcastKeyPressedData()
	}

	s.WG.Wait()
}

func (s *Server) listenDaemon() {
	// Listenning loop
	for {
		conn, err := s.Listener.Accept()
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

		switch gjson.Get(packet, "type").Int() {
		case protocol.PlayerJoinType:
			// TODO Don't accept player's join request after game starting
			newPlayerName := gjson.Get(packet, "data.name").String()
			s.Connections[newPlayerName] = conn
		case protocol.KeyPressType:
			s.KeyPressChan <- packet
		}
	}
}

func (s *Server) broadcastKeyPressedData() {
	defer s.WG.Done()

	lastTime := time.Now().UnixMilli()
	for {
		packet := <-s.KeyPressChan
		currentTime := time.Now().UnixMilli()
		if currentTime-lastTime >= 50 {
			lastTime = currentTime
			for _, conn := range s.Connections {
				conn.Write([]byte(packet))
			}
		}
	}
}

func (s *Server) broadcastSpaceships() {
	const START_TIME = 10
	defer s.WG.Done()

	for currentTime := 0; currentTime < START_TIME; currentTime++ {
		log.Printf("The game will start after %v s", START_TIME-currentTime)
		time.Sleep(time.Second)
	}

	// broadcast new player's space ship info
	playerList := []string{}
	for playName := range s.Connections {
		playerList = append(playerList, playName)
	}
	spaceShipData, err := protocol.GenerateSpaceShipPack(playerList)
	util.HandleErr(err)
	for _, conn := range s.Connections {
		conn.Write(spaceShipData)
	}
}

// TODO gracefully exit daemon for exit signal
