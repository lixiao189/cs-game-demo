package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lixiao189/cs-game-demo/client/game"
)

func main() {
	game := game.Game{
		Height: 640,
		Width:  1024,
	}

	ebiten.SetWindowSize(game.Width, game.Height)
	ebiten.SetWindowTitle("Space ship Demo!")

	game.InitGame()
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
