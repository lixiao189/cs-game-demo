package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lixiao189/cs-game-demo/protocol"
	"github.com/lixiao189/cs-game-demo/shape"
	"github.com/lixiao189/cs-game-demo/util"
	"github.com/tidwall/gjson"
)

func (g *Game) initDaemon() {

	g.WG.Add(2)
	{
		go g.receivePack()
		go g.handleKeyPressed()
	}
}

func (g *Game) receivePack() {
	defer g.WG.Done()
	for {
		buf := make([]byte, 1024)
		n, err := g.Conn.Read(buf)
		util.LogErr(err)
		pack := string(buf[:n])

		switch gjson.Get(pack, "type").Int() {
		case protocol.InitSpaceshipType:
			spaceshipList := gjson.Get(pack, "data").Array()
			go g.initSpaceship(spaceshipList)
		case protocol.KeyPressType:
			g.KeyPressedChan <- pack
		}
	}
}

func (g *Game) handleKeyPressed() {
	defer g.WG.Done()

	for {
		pack := <-g.KeyPressedChan

		keyPressed := ebiten.Key(gjson.Get(pack, "data.key").Int())
		playerName := gjson.Get(pack, "data.name").String()

		if spaceShip := g.SpaceShips[playerName]; spaceShip != nil {
			switch keyPressed {
			case ebiten.KeyW:
				spaceShip.Direction = shape.UP
				spaceShip.Y -= spaceShip.Speed
			case ebiten.KeyA:
				spaceShip.Direction = shape.LEFT
				spaceShip.X -= spaceShip.Speed
			case ebiten.KeyS:
				spaceShip.Direction = shape.DOWN
				spaceShip.Y += spaceShip.Speed
			case ebiten.KeyD:
				spaceShip.Direction = shape.RIGHT
				spaceShip.X += spaceShip.Speed
			}
		}
	}
}

func (g *Game) initSpaceship(spaceshipList []gjson.Result) {
	for index := range spaceshipList {
		name := spaceshipList[index].Get("name").String()
		x := spaceshipList[index].Get("x").Float()
		y := spaceshipList[index].Get("y").Float()

		g.SpaceShips[name] = shape.NewSpaceShip(x, y, 3, 64, 32, name)
	}
}
