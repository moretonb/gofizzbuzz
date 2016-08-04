package main

import "fmt"
import "gofizzbuzz/game"

func main() {
	for index := 1; index <= 100; index++ {
		fmt.Println(game.Play(index))
	}
}
