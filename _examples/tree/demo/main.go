package main

import (
	"github.com/mochen302/pterm/putils"
	"github.com/pterm/pterm"
)

func main() {
	// You can use a LeveledList here, for easy generation.
	leveledList := pterm.LeveledList{
		pterm.LeveledListItem{Level: 0, Text: "C:"},
		pterm.LeveledListItem{Level: 1, Text: "Users"},
		pterm.LeveledListItem{Level: 1, Text: "Windows"},
		pterm.LeveledListItem{Level: 1, Text: "Programs"},
		pterm.LeveledListItem{Level: 1, Text: "Programs(x86)"},
		pterm.LeveledListItem{Level: 1, Text: "dev"},
		pterm.LeveledListItem{Level: 0, Text: "D:"},
		pterm.LeveledListItem{Level: 0, Text: "E:"},
		pterm.LeveledListItem{Level: 1, Text: "Movies"},
		pterm.LeveledListItem{Level: 1, Text: "Music"},
		pterm.LeveledListItem{Level: 2, Text: "LinkinPark"},
		pterm.LeveledListItem{Level: 1, Text: "Games"},
		pterm.LeveledListItem{Level: 2, Text: "Shooter"},
		pterm.LeveledListItem{Level: 3, Text: "CallOfDuty"},
		pterm.LeveledListItem{Level: 3, Text: "CS:GO"},
		pterm.LeveledListItem{Level: 3, Text: "Battlefield"},
		pterm.LeveledListItem{Level: 4, Text: "Battlefield 1"},
		pterm.LeveledListItem{Level: 4, Text: "Battlefield 2"},
		pterm.LeveledListItem{Level: 0, Text: "F:"},
		pterm.LeveledListItem{Level: 1, Text: "dev"},
		pterm.LeveledListItem{Level: 2, Text: "dops"},
		pterm.LeveledListItem{Level: 2, Text: "PTerm"},
	}

	// Generate tree from LeveledList.
	root := putils.TreeFromLeveledList(leveledList)

	// Render TreePrinter
	pterm.DefaultTree.WithRoot(root).Render()
}
