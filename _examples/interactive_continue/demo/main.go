package main

import (
	"github.com/mochen302/pterm"
)

func main() {
	result, _ := pterm.DefaultInteractiveContinue.Show()
	pterm.Println() // Blank line
	pterm.Info.Printfln("You answered: %s", result)
}
