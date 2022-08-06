package protocol

import (
	"encoding/json"
	"math/rand"
)

func GeneratePack(packType string, data any) ([]byte, error) {
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

func GenerateSpaceShipPack(playerList []string) ([]byte, error) {
	spaceshipList := []map[string]any{}
	for index := range playerList {
		spaceshipList = append(spaceshipList, map[string]any{
			"x":    float64(rand.Intn(500)),
			"y":    float64(rand.Intn(400)),
			"name": playerList[index],
		})
	}
	return GeneratePack(InitSpaceshipType, spaceshipList)
}
