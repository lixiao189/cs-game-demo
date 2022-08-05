package protocol

import "github.com/lixiao189/cs-game-demo/shape"

type Pack struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

/*  ============================================================= */

const (
	// client data type
	PlayerJoinType = "join"
)

type JoinData struct {
	Name string `json:"name"`
}

/*  ============================================================= */

const (
	// server data type
	InitSpaceshipType = "init-spaceship"
)

type SpaceshipData struct {
	shape.Spaceship
}
