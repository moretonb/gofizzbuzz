package coordinator

import (
	"gofizzbuzz/game"
	"gofizzbuzz/output"
)

func CreateMaster() *Master {
	return newMaster(game.CreateGameRunner, output.CreatePrinter)
}
