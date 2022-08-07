package main

import (
	"net"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/lixiao189/cs-game-demo/game"
	"github.com/lixiao189/cs-game-demo/server"
)

var opts struct {
	Deamon bool   `short:"d" long:"deamon" description:"Whether working on deamon server"`
	Name   string `short:"n" long:"name" description:"Username of the player"`
}

func main() {
	// Parse flags
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(0)
	}

	if opts.Deamon {
		// Server initialize
		s := server.Server{
			Host: "0.0.0.0",
			Port: 1234,
			Connections: make(map[string]net.Conn),
		}
		s.ServerInit()
	} else {
		// Client initialize and startup code
		game := game.Game{
			Height: 640,
			Width:  1024,
			Host:   "127.0.0.1",
			Port:   1234,
		}
		game.InitGame(opts.Name)
	}
}
