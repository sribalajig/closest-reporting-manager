package interactions

import (
	"closest-reporting-manager/display"
	"closest-reporting-manager/models"
	"fmt"
)

func FindClosestManager(org models.Org) {
	display.ClearScreen()

	empOne := display.Input("Please enter the first employee.")

	found := org.Search(org.CEO, empOne)

	if found == nil {
		display.Prompt(fmt.Sprintf("You entered '%s', this person does not exist in the org. Press ENTER for the main menu", empOne))

		return
	}

	empTwo := display.Input("Please enter the second employee.")

	found = org.Search(org.CEO, empTwo)

	if found == nil {
		display.Prompt(fmt.Sprintf("You entered '%s', this person does not exist in the org. Press ENTER for the main menu", empTwo))

		return
	}

	commonManager := org.FindClosestCommonManager(empOne, empTwo)

	if commonManager == nil {
		display.Inform("Something went wrong.. :(")
	} else {
		display.Inform(fmt.Sprintf("The closest reporting manager between '%s' and '%s' is '%s'", empOne, empTwo, commonManager.Name))
	}

	display.Prompt("Press ENTER to return to the main menu")
}
