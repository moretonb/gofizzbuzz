package output

import (
	"os"
	"testing"
)

var printer = newPrinter(mockOutputPrinter)
var printedItems = make([]string, 3)
var currentPrintedIndex = 0
var inputChannel = make(chan string)
var complete = make(chan struct{})
var input = []string{"1", "2", "badger"}

func TestMain(m *testing.M) {
	go steamInputToPrint()
	go printer.Print(inputChannel, complete)
	<-complete

	result := m.Run()
	os.Exit(result)
}

func TestWhenPrintThenAllItemsArePrintedInTheExpectedOrder(t *testing.T) {
	t.Parallel()

	for index, result := range printedItems {
		if result != input[index] {
			t.Log("expected output ", input[index])
			t.Log("actual output ", result)
			t.Fail()
		}
	}
}

func steamInputToPrint() {
	for _, item := range input {
		inputChannel <- item
	}

	close(inputChannel)
}

func mockOutputPrinter(input string) {
	printedItems[currentPrintedIndex] = input
	currentPrintedIndex++
}
