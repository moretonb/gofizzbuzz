package game

import (
	"os"
	"testing"
)

var runner = newGameRunner(mockTurnTaker)
var resultsChannel = make(chan string)
var results = make([]string, 100)
var resultsLength = 0
var expectedTurnOutput = "badger"
var expectedResultsLength = 100

func TestMain(m *testing.M) {
	go runner.Run(resultsChannel)

	for result := range resultsChannel {
		results[resultsLength] = result
		resultsLength++
	}

	result := m.Run()
	os.Exit(result)
}

func TestWhenRunThenAllResultsHaveTheExpectedString(t *testing.T) {
	t.Parallel()

	for _, result := range results {
		if result != expectedTurnOutput {
			t.Log("expected output ", expectedTurnOutput)
			t.Log("actual output ", result)
			t.Fail()
		}
	}
}

func TestWhenRunThenTakeTurnIsCalled100Times(t *testing.T) {
	t.Parallel()

	if resultsLength != expectedResultsLength {
		t.Log("expected output ", expectedResultsLength)
		t.Log("actual output ", resultsLength)
		t.Fail()
	}
}

func mockTurnTaker(input int) string {
	return expectedTurnOutput
}
