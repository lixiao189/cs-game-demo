package server

// TODO Closing connection gracefully

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/lixiao189/cs-game-demo/protocol"
	"github.com/lixiao189/cs-game-demo/util"
	"github.com/xtaci/kcp-go/v5"
	"golang.org/x/exp/rand"
)

func ServerInit(host string, port int) {
	log.Println("Running on server mode")

	laddr := fmt.Sprintf("%v:%v", host, port)
	listener, err := kcp.Listen(laddr)
	if err != nil {
		defer listener.Close()
		log.Fatal(err)
		os.Exit(-1)
	}

	for {
		serverConn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handlePacket(serverConn)
	}
}

func handlePacket(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		util.HandleErr(err)
		log.Println(string(buf))

		var clientPack protocol.Pack
		err = json.Unmarshal(buf[:n], &clientPack)
		util.HandleErr(err)

		switch clientPack.Type {
		case protocol.PlayerJoinType:
			newPlayerName := clientPack.Data.(map[string]interface{})["name"].(string)
			log.Println(newPlayerName + " joins the game")

			spaceshipData, err := json.Marshal(protocol.Pack{
				Type: protocol.InitSpaceshipType,
				Data: protocol.SpaceshipData{
					X: float64(rand.Intn(500)),
					Y: float64(rand.Intn(300)),
					Speed: 3,
					Height: 64,
					Width: 32,
					Name: newPlayerName,
				},
			})
			util.HandleErr(err)
			conn.Write(spaceshipData)
		}
	}
}
