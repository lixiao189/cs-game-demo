package game

// TODO Closing connection gracefully

import (
	"fmt"
	"log"
	"math"
	"net"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lixiao189/cs-game-demo/client/shape"
	"github.com/xtaci/kcp-go/v5"
)

type Game struct {
	SpaceShips map[string]*shape.Spaceship
	PlayerName string

	// Window size
	Height int
	Width  int

	// Server host info
	Host       string
	Port       int
	ClientConn net.Conn
}

func (g *Game) Update() error {
	playerSpaceShip := g.SpaceShips[g.PlayerName]
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		playerSpaceShip.Direction = shape.LEFT
		playerSpaceShip.X -= playerSpaceShip.Speed
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		playerSpaceShip.Direction = shape.RIGHT
		playerSpaceShip.X += playerSpaceShip.Speed
	} else if ebiten.IsKeyPressed(ebiten.KeyW) {
		playerSpaceShip.Direction = shape.UP
		playerSpaceShip.Y -= playerSpaceShip.Speed
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		playerSpaceShip.Direction = shape.DOWN
		playerSpaceShip.Y += playerSpaceShip.Speed
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// debug
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%v\n%v %v", g.PlayerName, g.SpaceShips[g.PlayerName].X, g.SpaceShips[g.PlayerName].Y))

	for _, spaceShip := range g.SpaceShips {
		op := &ebiten.DrawImageOptions{}

		// Rotate the space ship
		op.GeoM.Reset()
		op.GeoM.Translate(-float64(spaceShip.Width)/2, -float64(spaceShip.Height)/2)
		if spaceShip.Direction == shape.LEFT {
			op.GeoM.Rotate(math.Pi * 3 / 2)
		} else if spaceShip.Direction == shape.RIGHT {
			op.GeoM.Rotate(math.Pi / 2)
		} else if spaceShip.Direction == shape.UP {
			op.GeoM.Rotate(0)
		} else if spaceShip.Direction == shape.DOWN {
			op.GeoM.Rotate(math.Pi)
		}

		// Move to the right place
		op.GeoM.Translate(spaceShip.X, spaceShip.Y)

		screen.DrawImage(spaceShip.Image, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}

func (g *Game) InitGame(name string) {
	log.Println("Game started")

	// Player's name
	playerName := name

	// Init game system
	g.PlayerName = playerName
	g.SpaceShips = make(map[string]*shape.Spaceship)
	g.SpaceShips[playerName] = shape.NewSpaceShip(
		float64(g.Width)/2,
		float64(g.Height)/2,
		3, 64, 32,
		playerName,
	)

	// Connect to server host
	raddr := fmt.Sprintf("%v:%v", g.Host, g.Port)
	clientConn, err := kcp.Dial(raddr)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
	g.ClientConn = clientConn

	// Init ebiten window's setting
	ebiten.SetWindowSize(g.Width, g.Height)
	ebiten.SetWindowTitle("Space ship Demo!")

	// Running game
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
