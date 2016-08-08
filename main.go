package main

import (
	"gofizzbuzz/coordinator"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)

	master := coordinator.CreateMaster()
	master.Start()
}
