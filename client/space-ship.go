package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lixiao189/cs-game-demo/client/resources"
)

type Spaceship struct {
	Image *ebiten.Image

	// Space ship position
	X int
	Y int

	// Space ship size
	Height int
	Width  int

	// Master's name
	Name string
}

func NewSpaceShip(x int, y int, height int, width int, name string) *Spaceship {
	spaceShip := Spaceship{
		X:      x,
		Y:      y,
		Height: height,
		Width:  width,
		Name:   name,
	}
	img, _, err := image.Decode(bytes.NewReader(resources.SpaceshipImg))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	spaceShip.Image = ebiten.NewImageFromImage(img)
	return &spaceShip
}
