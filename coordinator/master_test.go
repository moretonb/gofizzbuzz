package coordinator

import (
	"gofizzbuzz/game"
	"gofizzbuzz/output"
	"os"
	"testing"
)

var master = newMaster(mockGameRunnerCreator, mockPrinterCreator)
var runnerCalled = false
var printerCalled = false

func TestMain(m *testing.M) {
	master.Start()

	result := m.Run()
	os.Exit(result)
}

func TestWhenStartThenGameRunnerIsCalled(t *testing.T) {
	t.Parallel()

	if !runnerCalled {
		t.Error("Expected runner to be called")
	}
}

func TestWhenStartThenPrinterIsCalled(t *testing.T) {
	t.Parallel()

	if !printerCalled {
		t.Error("Expected printer to be called")
	}
}

func mockGameRunnerCreator() game.IGameRunner {
	return &MockGameRunner{}
}

type MockGameRunner struct {
}

func (mgr *MockGameRunner) Run(results chan string) {
	runnerCalled = true
}

func mockPrinterCreator() output.IPrinter {
	return &MockPrinter{}
}

type MockPrinter struct {
}

func (mo *MockPrinter) Print(results chan string, complete chan bool) {
	printerCalled = true

	close(complete)
}
