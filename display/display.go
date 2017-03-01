package display

import (
	"bufio"
	"fmt"
	"os"
)

func AskQuestion(question string) {

}

func Prompt(prompt string) rune {
	fmt.Println()

	fmt.Println(prompt)

	fmt.Println()

	return GetChar()
}

func Input(prompt string) string {
	fmt.Println()

	fmt.Println(prompt)

	fmt.Println()

	var input string

	for {
		input = GetString()

		if input != "" {
			break
		} else {
			fmt.Println()

			fmt.Println("Looks like you entered a blank name. That's invalid!")

			fmt.Println()

			fmt.Println(prompt)
		}
	}

	return input
}

func Inform(information string) {
	fmt.Println()

	fmt.Println(information)

	fmt.Println()
}

func ShowMenu(menu []string) {
	fmt.Println()

	fmt.Println("..............................................................")

	fmt.Println()

	for _, menuItem := range menu {
		fmt.Println()
		fmt.Println(menuItem)
	}

	fmt.Println()

	fmt.Println("..............................................................")
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func GetChar() rune {
	reader := bufio.NewReader(os.Stdin)

	ch, _, _ := reader.ReadRune()

	return ch
}

func GetString() string {
	reader := bufio.NewReader(os.Stdin)

	input, _, _ := reader.ReadLine()

	return string(input)
}
