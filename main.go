package main

import (
	"closest-reporting-manager/display"
	"closest-reporting-manager/interactions"
	"closest-reporting-manager/models"
	"closest-reporting-manager/org-service"
	"fmt"
)

var org models.Org

func init() {
	org = service.Create("bureaucr.at")
}

func main() {
	display.ClearScreen()

	display.Inform(fmt.Sprintf("Welcome! This program lets manage the org tree for %s", org.Name))

	display.Prompt("Press ENTER to get started")

	display.ClearScreen()

	if org.CEO == nil {
		display.Prompt(fmt.Sprintf("Looks like %s does not have a CEO. Press ENTER to add your CEO.", org.Name))

		name := display.Input("Whats your CEO's name?")

		ceo := models.Employee{
			Name: name,
		}

		org.CEO = &ceo

		display.Inform(fmt.Sprintf("Great! %s was added the CEO of %s.", ceo.Name, org.Name))

		display.Prompt("Press ENTER for the menu")
	}

	interactions.Menu(org)
}
