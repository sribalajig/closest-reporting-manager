package main

import (
	"bufio"
	"closest-reporting-manager/display"
	"closest-reporting-manager/models"
	"closest-reporting-manager/org-service"
	"fmt"
	"os"
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
		display.Prompt(fmt.Sprintf("Looks like %s does not have a CEO. Press ENTER key to add your CEO.", org.Name))

		fmt.Println("Whats your CEO's name?")

		ceo := models.Employee{
			Name: string(getString()),
		}

		org.CEO = &ceo

		display.Inform(fmt.Sprintf("Great! %s was added the CEO of %s.", ceo.Name, org.Name))

		display.Prompt("Press ENTER for the menu")
	}

	for {
		display.ClearScreen()

		display.Inform(fmt.Sprintf("This is is the org structure for %s", org.Name))

		org.PrintLevels()

		display.ShowMenu([]string{
			"\tTo add a person to the org, PRESS 1",
			"\tTo find closest reporting manager, PRESS 2",
		})

		option := display.GetChar()

		if string(option) == "1" {
			AddEmployee()
		}

		if string(option) == "2" {
			FindClosestManager()
		}
	}
}

func AddEmployee() {
	display.ClearScreen()

	newEmp := display.Input("What's the employees name?")

	existing := org.Search(org.CEO, newEmp)

	if existing != nil {
		display.Prompt(fmt.Sprintf("You entered '%s'. This person already exists in the org.", existing.Name))

		return
	}

	manager := display.Input("Whom does this person report to")

	existing = org.Search(org.CEO, manager)

	if existing == nil {
		display.Prompt(fmt.Sprintf("You entered '%s' as the reporting manager for '%s'.There is no such employee in the org", manager, newEmp))

		return
	}

	existing.AddReport(newEmp)

	display.Inform(fmt.Sprintf("'%s' was added the org!", newEmp))

	display.Prompt("Press ENTER to return to the main menu")
}

func FindClosestManager() {
	fmt.Print("\033[H\033[2J")

	fmt.Println("Enter the first employee")

	empOne := getString()

	found := org.Search(org.CEO, empOne)

	if found == nil {
		fmt.Println(fmt.Sprintf("You entered '%s', this person does not exist in the org", empOne))

		getChar()

		return
	}

	fmt.Println("Enter the second employee")

	empTwo := getString()

	found = org.Search(org.CEO, empTwo)

	if found == nil {
		fmt.Println(fmt.Sprintf("You entered '%s', this person does not exist in the org", empTwo))

		getChar()

		return
	}

	commonManager := org.FindClosestCommonManager(empOne, empTwo)

	fmt.Println()

	if commonManager == nil {
		fmt.Println("Something went wrong.. :(")
	} else {
		fmt.Println(fmt.Sprintf("The closest reporting manager between '%s' and '%s' is '%s'", empOne, empTwo, commonManager.Name))
	}

	getChar()
}

func getChar() rune {
	reader := bufio.NewReader(os.Stdin)

	ch, _, _ := reader.ReadRune()

	return ch
}

func getString() string {
	reader := bufio.NewReader(os.Stdin)

	input, _, _ := reader.ReadLine()

	return string(input)
}
