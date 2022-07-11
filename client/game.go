package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	SpaceShips map[string]*Spaceship

	// Window size
	Height int
	Width  int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, spaceShip := range(g.SpaceShips) {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(spaceShip.X), float64(spaceShip.Y))
		screen.DrawImage(spaceShip.Image, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}

func (g *Game) InitGame() {
	log.Println("Game started")

	// Player's name
	playerName := "node"
	g.SpaceShips = make(map[string]*Spaceship)
	g.SpaceShips[playerName] = NewSpaceShip(g.Width/2, g.Height/2, 64, 32, playerName)
}
