package game

type turnTaker func(input int) string

type IGameRunner interface {
	Run(chan string)
}

type GameRunner struct {
	takeTurn turnTaker
}

func newGameRunner(tt turnTaker) *GameRunner {
	return &GameRunner{takeTurn: tt}
}

func (gr *GameRunner) Run(results chan string) {
	for index := 1; index <= 100; index++ {
		results <- gr.takeTurn(index)
	}

	close(results)
}
