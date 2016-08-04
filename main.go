package main

import (
	"fmt"
	"gofizzbuzz/game"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)

	gameResults := make(chan string)
	complete := make(chan bool)

	go play1to100(gameResults)
	go printer(gameResults, complete)

	<-complete
}

func play1to100(gameResults chan string) {
	for index := 1; index <= 100; index++ {
		gameResults <- game.Play(index)
	}

	close(gameResults)
}

func printer(gameResults chan string, complete chan bool) {
	for result := range gameResults {
		fmt.Println(result)
	}

	complete <- true
}
