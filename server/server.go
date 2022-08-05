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

func handlePacket(serverConn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := serverConn.Read(buf)
		util.HandleErr(err)
		log.Println(string(buf))

		var clientPack protocol.Pack
		err = json.Unmarshal(buf[:n], &clientPack)
		util.HandleErr(err)
		
		switch clientPack.Type {
		case protocol.PlayerJoinType:
			newPlayerName := clientPack.Data.(map[string]interface{})["name"]
			log.Println(newPlayerName)
		}
	}
}
