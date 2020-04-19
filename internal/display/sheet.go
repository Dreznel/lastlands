package display

// Goal: I want to make a method/function/algorithm/object that can print
// character sheets in a neat, text-based manner. It should have multiple columns
// per row with potentially unrelated content.

// If this proves too difficult, I might just scrap this, but I thought it'd be fun
// to try.

// TODO: try https://github.com/olekukonko/tablewriter

import (
	"github.com/dreznel/lastlands/v2/internal/character"
	"github.com/rodaine/table"
)

func GetCharacterSheetTables(seeker character.Seeker) table.Table {
	tbl := table.New("CORE Memories", "Stats")
	tbl.AddRow(CoreMemoriesTable(seeker), BasicStatTable(seeker))
	return tbl
}

func BasicStatTable(seeker character.Seeker) table.Table {
	tbl := table.New("Attribute", "Bonus", "Defense")
	stats := character.AbilityScoresToMap(seeker.AbilityScores)
	for key, abilityScore := range stats {
		tbl.AddRow(key, abilityScore.Bonus, abilityScore.Defense)
	}

	return tbl
}

func CoreMemoriesTable(seeker character.Seeker) table.Table {
	tbl := table.New("d6", "Type", "Memory")
	tbl.AddRow("1/2", seeker.CoreAttributes.First.GetType(), seeker.CoreAttributes.First.GetDescription())
	tbl.AddRow("3/4", seeker.CoreAttributes.Second.GetType(), seeker.CoreAttributes.Second.GetDescription())
	tbl.AddRow("5/6", seeker.CoreAttributes.Third.GetType(), seeker.CoreAttributes.Third.GetDescription())
	return tbl
}

// func example() {
// 	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
// 	columnFmt := color.New(color.FgYellow).SprintfFunc()

// 	tbl := table.New("ID", "Name", "Score", "Added")
// 	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

// 	for _, widget := range getWidgets() {
// 		tbl.AddRow(widget.ID, widget.Name, widget.Cost, widget.Added)
// 	}

// 	tbl.Print()
// }
