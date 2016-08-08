package output

import "fmt"

func CreatePrinter() IPrinter {
	return newPrinter(printLnWrapper)
}

func printLnWrapper(input string) {
	fmt.Println(input)
}
