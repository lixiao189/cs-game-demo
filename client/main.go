package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := Game{
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
