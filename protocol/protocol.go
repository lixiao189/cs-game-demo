package protocol

import "encoding/json"

func GeneratePack(packType string, data map[string]any) ([]byte, error) {
	packBytes, err := json.Marshal(map[string]any{
		"type": packType,
		"data": data,
	})

	return packBytes, err
}

/*  ============================================================= */

const (
	// client data type
	PlayerJoinType = "join"
)

func GenerateJoinPack(playerName string) ([]byte, error) {
	return GeneratePack(PlayerJoinType, map[string]any{
		"name": playerName,
	})
}

/*  ============================================================= */

const (
	// server data type
	InitSpaceshipType = "init-spaceship"
)

