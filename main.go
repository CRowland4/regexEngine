package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	regex, input := parseInput()
	fmt.Println(isRegexMatch(regex, input))
}

// From stage three.
func isRegexMatch(regex, input string) (isMatch bool) {
	if isBaseCaseMet, result := isBaseCase(regex, input); isBaseCaseMet {
		return result == "true"
	}

	if strings.HasPrefix(regex, "^") {
		return isEqualLengthRegexMatch(regex[1:], input)
	}
	if isEqualLengthRegexMatch(regex, input) {
		return true
	}

	return isRegexMatch(regex, input[1:])
}

// From stage two.
func isEqualLengthRegexMatch(regex, input string) (isMatch bool) {
	if isBaseCaseMet, result := isBaseCase(regex, input); isBaseCaseMet {
		return result == "true"
	}

	if strings.Index(regex, "?") == 1 {
		return handleQuestionMark(regex, input)
	}
	if strings.Index(regex, "*") == 1 {
		return handleAsterisk(regex, input)
	}

	if strings.Index(regex, "+") == 1 {
		return handlePlus(regex, input, false)
	}

	if !isCharacterMatch(regex[0], input[0]) {
		return false
	}

	return isEqualLengthRegexMatch(regex[1:], input[1:])
}

func handlePlus(regex, input string, plusFlag bool) (isMatch bool) {
	if isBaseCaseMet, result := isBaseCase(regex, input); isBaseCaseMet {
		return result == "true"
	}
	if !isCharacterMatch(regex[0], input[0]) && !plusFlag {
		return false
	}

	if !isCharacterMatch(regex[0], input[0]) && plusFlag {
		regex = strings.Replace(regex, regex[0:2], "", 1)
		return isEqualLengthRegexMatch(regex, input)
	}

	// Characters match
	return handlePlus(regex, input[1:], true)

}

func handleAsterisk(regex, input string) (isMatch bool) {
	if !isCharacterMatch(regex[0], input[0]) {
		regex = strings.Replace(regex, regex[0:2], "", 1)
		return isEqualLengthRegexMatch(regex, input)
	}

	return isEqualLengthRegexMatch(regex, input[1:])
}

func handleQuestionMark(regex, input string) (isMatch bool) {
	if isCharacterMatch(regex[0], input[0]) {
		input = input[1:]
	}

	regex = strings.Replace(regex, regex[0:2], "", 1)
	return isEqualLengthRegexMatch(regex, input)
}

func isBaseCase(regex, input string) (isMet bool, result string) {
	if (regex == ".*") && (input == "") {
		return true, "true"
	}
	if (regex == ".+") && (input == "") {
		return true, "true"
	}
	if regex == "" {
		return true, "true"
	}
	if (regex == "$") && (input == "") {
		return true, "true"
	}
	if input == "" {
		return true, "false"
	}

	return false, "base case not met"
}

// From stage one. Because of base cases being checked before this function is called, neither argument will ever be ""
func isCharacterMatch(regexChar, inputChar byte) (isMatch bool) {
	return (regexChar == inputChar) || (string(regexChar) == ".")
}

func parseInput() (regex, input string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fullInput := scanner.Text()
	splitInput := strings.Split(fullInput, "|")

	return splitInput[0], splitInput[1]
}
