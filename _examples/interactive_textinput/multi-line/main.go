package main

import (
	"github.com/mochen302/pterm"
)

func main() {
	result, _ := pterm.DefaultInteractiveTextInput.WithMultiLine().Show() // Text input with multi line enabled
	pterm.Println()                                                       // Blank line
	pterm.Info.Printfln("You answered: %s", result)
}
