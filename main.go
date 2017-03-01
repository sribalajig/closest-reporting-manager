package main

import (
	"bufio"
	"closest-reporting-manager/models"
	"closest-reporting-manager/services"
	"fmt"
	"os"
)

var org models.Org

func init() {
	org = models.Org{
		Name: "bureaucr.at",
	}
}

func main() {
	fmt.Print("\033[H\033[2J")

	fmt.Println(fmt.Sprintf("Welcome! This program lets manage the org tree for %s", org.Name))

	fmt.Println()

	fmt.Println(fmt.Sprintf("Press ENTER key to get started"))

	getChar()

	fmt.Print("\033[H\033[2J")

	if org.CEO == nil {
		fmt.Println(fmt.Sprintf("Looks like %s does not have a CEO. Press ENTER key to add your CEO.", org.Name))

		getChar()

		fmt.Println("Whats your CEO's name?")

		ceo := models.Employee{
			Name: string(getString()),
		}

		org.CEO = &ceo

		fmt.Println()

		fmt.Println(fmt.Sprintf("Great! %s was added the CEO of %s.", ceo.Name, org.Name))
	}

	for {
		fmt.Print("\033[H\033[2J")

		fmt.Println(fmt.Sprintf("This is is the org structure for %s", org.Name))

		services.Bfs(org.CEO)

		fmt.Println()
		fmt.Println("What would you like to do now?")
		fmt.Println()

		fmt.Println("\tTo add a person to the org, PRESS 1")
		fmt.Println()
		fmt.Println("\tTo find closest reporting manager, PRESS 2")

		fmt.Println()

		option := getChar()

		if string(option) == "1" {
			AddEmployee()
		}

		if string(option) == "2" {
			FindClosestManager()
		}
	}
}

func AddEmployee() {
	fmt.Print("\033[H\033[2J")

	fmt.Println("Whats the employees name")

	newEmp := getString()

	existing := services.Dfs(org.CEO, newEmp)

	if existing != nil {
		fmt.Println(fmt.Sprintf("You entered '%s'. This person already exists in the org.", existing.Name))

		getChar()

		return
	}

	fmt.Println()

	fmt.Println("Whom does this person report to")

	manager := getString()

	existing = services.Dfs(org.CEO, manager)

	if existing == nil {
		fmt.Println(fmt.Sprintf("You entered '%s' as the reporting manager for '%s'.There is no such employee in the org", manager, newEmp))

		getChar()

		return
	}

	existing.AddReport(newEmp)

	fmt.Println()

	fmt.Println(fmt.Sprintf("'%s' was added the org!", newEmp))

	getChar()
}

func FindClosestManager() {
	fmt.Print("\033[H\033[2J")

	fmt.Println("Enter the first employee")

	empOne := getString()

	found := services.Dfs(org.CEO, empOne)

	if found == nil {
		fmt.Println(fmt.Sprintf("You entered '%s', this person does not exist in the org", empOne))

		getChar()

		return
	}

	fmt.Println("Enter the second employee")

	empTwo := getString()

	found = services.Dfs(org.CEO, empTwo)

	if found == nil {
		fmt.Println(fmt.Sprintf("You entered '%s', this person does not exist in the org", empTwo))

		getChar()

		return
	}

	commonManager := services.FindLeastCommonAncestor(org.CEO, empOne, empTwo)

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
