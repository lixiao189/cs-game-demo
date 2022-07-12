package main

import (
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/lixiao189/cs-game-demo/client/game"
	"github.com/lixiao189/cs-game-demo/server"
)

var opts struct {
	// Slice of bool will append 'true' each time the option
	// is encountered (can be set multiple times, like -vvv)
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
		server.ServerInit("127.0.0.1", 1234)
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
