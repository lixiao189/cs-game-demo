package game

import (
	"log"

	"github.com/lixiao189/cs-game-demo/protocol"
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

		switch gjson.Get(pack, "type").String() {
		case protocol.InitSpaceshipType:
			log.Println(pack) // debug
		}
	}
}
