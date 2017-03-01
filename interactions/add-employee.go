package interactions

import (
	"closest-reporting-manager/display"
	"closest-reporting-manager/models"
	"fmt"
)

func AddEmployee(org models.Org) {
	display.ClearScreen()

	newEmp := display.Input("What's the employees name?")

	existing := org.Search(org.CEO, newEmp)

	if existing != nil {
		display.Prompt(fmt.Sprintf("You entered '%s'. This person already exists in the org. Press ENTER for the main menu", existing.Name))
		return
	}

	manager := display.Input("Whom does this person report to")

	existing = org.Search(org.CEO, manager)

	if existing == nil {
		display.Prompt(fmt.Sprintf("You entered '%s' as the reporting manager for '%s'.There is no such employee in the org. Press ENTER for the main menu", manager, newEmp))

		return
	}

	existing.AddReport(newEmp)

	display.Inform(fmt.Sprintf("'%s' was added the org!", newEmp))

	display.Prompt("Press ENTER to return to the main menu")
}
