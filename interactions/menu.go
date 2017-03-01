package interactions

import (
	"closest-reporting-manager/display"
	"closest-reporting-manager/models"
	"fmt"
)

func Menu(org models.Org) {
	for {
		display.ClearScreen()

		display.Inform(fmt.Sprintf("Org hierarchy for : %s", org.Name))

		org.PrintLevels()

		display.ShowMenu([]string{
			"\tTo add a person to the org, PRESS 1",
			"\tTo find closest reporting manager, PRESS 2",
			"\tTo find closest reporting manager, PRESS e",
		})

		option := display.GetChar()

		if string(option) == "1" {
			AddEmployee(org)
		}

		if string(option) == "2" {
			FindClosestManager(org)
		}

		if string(option) == "e" {
			return
		}
	}
}
