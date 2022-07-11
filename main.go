package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jessevdk/go-flags"
	"github.com/lixiao189/cs-game-demo/client/game"
)

var opts struct {
	// Slice of bool will append 'true' each time the option
	// is encountered (can be set multiple times, like -vvv)
	Deamon bool `short:"d" long:"deamon" description:"Whether working on deamon server"`
}

func main() {
	// Parse flags
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(0)
	}

	if opts.Deamon {
		log.Println("Running on server mode")
		
	} else {
		// Client initialize and startup code
		game := game.Game{
			Height: 640,
			Width:  1024,
		}
		game.InitGame()
		if err := ebiten.RunGame(&game); err != nil {
			log.Fatal(err)
		}
	}
}
