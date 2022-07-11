package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	SpaceShips map[string]*Spaceship
	PlayerName string

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
		op.GeoM.Translate(spaceShip.X, spaceShip.Y)
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

	// Init game system
	g.PlayerName = playerName
	g.SpaceShips = make(map[string]*Spaceship)
	g.SpaceShips[playerName] = NewSpaceShip(float64(g.Width)/2, float64(g.Height)/2, 64, 32, playerName)
}
