package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	regex, input := parseInput()
	fmt.Println(regexMatch(regex, input))

	return
}

func regexMatch(regex, input string) bool {
	if regex == "" {
		return true
	}
	if input == "" {
		return false
	}
	if doesFirstCharacterMatch(regex, input) {
		return regexMatch(regex[1:], input[1:])
	}

	return regexMatch(regex, input[1:])
}

func doesFirstCharacterMatch(regex, input string) (isMatch bool) {
	return (string(regex[0]) == ".") || (regex[0] == input[0])
}

func parseInput() (regex, input string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fullInput := scanner.Text()
	splitInput := strings.Split(fullInput, "|")

	regex, input = splitInput[0], splitInput[1]
	return splitInput[0], splitInput[1]
}
