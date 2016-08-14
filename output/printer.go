package output

type outputPrinter func(input string)

type IPrinter interface {
	Print(chan string, chan struct{})
}

type Printer struct {
	printOutput outputPrinter
}

func newPrinter(op outputPrinter) *Printer {
	return &Printer{printOutput: op}
}

func (p *Printer) Print(results chan string, complete chan struct{}) {
	for result := range results {
		p.printOutput(result)
	}

	close(complete)
}
