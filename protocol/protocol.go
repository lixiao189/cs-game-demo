package protocol

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
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Speed  float64 `json:"speed"`
	Height int `json:"height"`
	Width  int `json:"width"`
	Name   string `json:"name"`
}
