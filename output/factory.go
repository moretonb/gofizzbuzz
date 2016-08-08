package output

import "fmt"

func CreatePrinter() IPrinter {
	return newPrinter(fmtWrapper)
}

func fmtWrapper(input string) {
	fmt.Println(input)
}
