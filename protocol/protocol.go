package protocol

const (
	// client data type
	PlayerJoinType = "join"
)

type ClientPack struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

type SpaceShipData struct {
	Name string `json:"name"`

	
}

/*  ============================================================= */
