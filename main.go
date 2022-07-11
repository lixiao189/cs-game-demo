package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lixiao189/cs-game-demo/client/game"
)

func main() {
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
