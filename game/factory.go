package game

func CreateGameRunner() IGameRunner {
	return newGameRunner(turn)
}
