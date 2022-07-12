package server

import (
	"fmt"
	"log"
	"net"
	"os"

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
	buf := make([]byte, 128)
	for {
		_, err := serverConn.Read(buf)

		if err != nil {
			log.Println(err)
			return
		} else {
			log.Println(string(buf))
		}
	}
}
