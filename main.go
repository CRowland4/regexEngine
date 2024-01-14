package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	regex, input := parseInput()

	if regex == input {
		fmt.Println("true")
	} else if regex == "." {
		fmt.Println("true")
	} else if regex == "" {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func parseInput() (regex, input string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fullInput := scanner.Text()
	splitInput := strings.Split(fullInput, "|")

	regex, input = splitInput[0], splitInput[1]
	return splitInput[0], splitInput[1]
}
