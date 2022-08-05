package protocol

const (
	// client data type
	PlayerJoinType = "join"
)

type ClientPack struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

type JoinData struct {
	Name string `json:"name"`
}

/*  ============================================================= */
