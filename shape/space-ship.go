package shape

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lixiao189/cs-game-demo/resources"
)

const (
	UP    = 0
	DOWN  = 1
	LEFT  = 2
	RIGHT = 3
)

type Spaceship struct {
	Image *ebiten.Image

	// Space ship theta
	Direction int

	// Space ship position
	X float64
	Y float64

	// Space ship speed
	Speed float64

	// Space ship size
	Height int
	Width  int

	// Master's name
	Name string
}

func NewSpaceShip(x float64, y float64, speed float64, height int, width int, name string) *Spaceship {
	spaceShip := Spaceship{
		X:         x,
		Y:         y,
		Direction: UP,
		Speed:     speed,
		Height:    height,
		Width:     width,
		Name:      name,
	}
	img, _, err := image.Decode(bytes.NewReader(resources.SpaceshipImg))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	spaceShip.Image = ebiten.NewImageFromImage(img)
	return &spaceShip
}
