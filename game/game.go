package game

// TODO Closing connection gracefully

import (
	"fmt"
	"log"
	"math"
	"net"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lixiao189/cs-game-demo/protocol"
	"github.com/lixiao189/cs-game-demo/shape"
	"github.com/lixiao189/cs-game-demo/util"
	"github.com/xtaci/kcp-go/v5"
)

type Game struct {
	SpaceShips map[string]*shape.Spaceship
	PlayerName string

	// Window size
	Height int
	Width  int

	// Server host info
	Host string
	Port int
	Conn net.Conn

	WG sync.WaitGroup
}

func (g *Game) Update() error {
	var keyPressed ebiten.Key
	isKeyPressed := true
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		keyPressed = ebiten.KeyA
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		keyPressed = ebiten.KeyD
	} else if ebiten.IsKeyPressed(ebiten.KeyW) {
		keyPressed = ebiten.KeyW
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		keyPressed = ebiten.KeyS
	} else {
		isKeyPressed = false
	}

	// Send the key pressed data to server
	if isKeyPressed {
		keyPressData, err := protocol.GenerateKeyPressPack(g.PlayerName, int(keyPressed))
		util.LogErr(err)
		g.Conn.Write(keyPressData)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	positionInfo := ""
	for _, spaceShip := range g.SpaceShips {
		positionInfo += fmt.Sprintf("%v\n%v %v\n", spaceShip.Name, spaceShip.X, spaceShip.Y)
	}
	ebitenutil.DebugPrint(screen, positionInfo)

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

	// Connect to server host and join the game
	clientConn, err := kcp.Dial(fmt.Sprintf("%v:%v", g.Host, g.Port))
	util.HandleErr(err)
	g.Conn = clientConn
	joinData, err := protocol.GenerateJoinPack(playerName)
	util.HandleErr(err)
	_, err = g.Conn.Write(joinData)
	util.HandleErr(err)

	// Init daemon gorountine after connecting
	g.initDaemon()

	// Init game system
	g.PlayerName = playerName
	g.SpaceShips = make(map[string]*shape.Spaceship)

	// Init ebiten window's setting
	ebiten.SetWindowSize(g.Width, g.Height)
	ebiten.SetWindowTitle("Space ship Demo!")

	// Running game
	err = ebiten.RunGame(g)
	util.HandleErr(err)

	g.WG.Wait()
}
