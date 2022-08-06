package game

import (
	"github.com/lixiao189/cs-game-demo/protocol"
	"github.com/lixiao189/cs-game-demo/shape"
	"github.com/lixiao189/cs-game-demo/util"
	"github.com/tidwall/gjson"
)

func (g *Game) initDaemon() {
	g.WG.Add(1)
	{
		go g.receivePack()
	}
}

func (g *Game) receivePack() {
	for {
		buf := make([]byte, 1024)
		n, err := g.Conn.Read(buf)
		util.LogErr(err)
		pack := string(buf[:n])

		switch gjson.Get(pack, "type").Int() {
		case protocol.InitSpaceshipType:
			spaceshipList := gjson.Get(pack, "data").Array()
			go g.initSpaceship(spaceshipList)
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
