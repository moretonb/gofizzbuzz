package coordinator

import (
	"gofizzbuzz/game"
	"gofizzbuzz/output"
)

type gameRunnerCreator func() game.IGameRunner
type printerCreator func() output.IPrinter

type Master struct {
	createGameRunner gameRunnerCreator
	createPrinter    printerCreator
}

func newMaster(grc gameRunnerCreator, pc printerCreator) *Master {
	return &Master{createGameRunner: grc, createPrinter: pc}
}

func (m *Master) Start() {
	gameResults := make(chan string)
	complete := make(chan bool)

	runner := m.createGameRunner()
	printer := m.createPrinter()

	go runner.Run(gameResults)
	go printer.Print(gameResults, complete)

	<-complete
}
