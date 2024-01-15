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

func isEqualLengthRegexMatch(regex, input string) (isMatch bool) {
	if isBaseCaseMet, result := isBaseCase(regex, input); isBaseCaseMet {
		return result == "true"
	}

	if (strings.Index(regex, "\\") == 0) && (strings.IndexAny(regex[1:], "^$?*+.\\")) == 0 {
		return handleEscapeCharacter(regex, input)
	}

	if strings.IndexAny(regex, "?*+") == 1 {
		return handleRepetitionOperators(regex, input)
	}

	if !isCharacterMatch(regex[0], input[0]) {
		return false
	}

	return isEqualLengthRegexMatch(regex[1:], input[1:])
}

func handleEscapeCharacter(regex, input string) (isMatch bool) {
	if !isRawCharacterMatch(regex[1], input[0]) {
		return isEqualLengthRegexMatch(regex, input[1:])
	}

	newRegex := strings.Replace(regex, regex[0:2], "", 1)
	return isEqualLengthRegexMatch(newRegex, input[1:])
}

func handleRepetitionOperators(regex, input string) (isMatch bool) {
	if strings.Index(regex, "?") == 1 {
		return handleQuestionMark(regex, input)
	}
	if strings.Index(regex, "*") == 1 {
		return handleAsterisk(regex, input)
	}

	return handlePlus(regex, input, false)
}

func handlePlus(regex, input string, plusFlag bool) (isMatch bool) {
	if isBaseCaseMet, result := isBaseCase(regex, input); isBaseCaseMet {
		return result == "true"
	}

	newRegex := strings.Replace(regex, regex[0:2], "", 1)
	if plusFlag && isEqualLengthRegexMatch(newRegex, input) {
		return true
	}

	if !isCharacterMatch(regex[0], input[0]) && !plusFlag {
		return false
	}
	if !isCharacterMatch(regex[0], input[0]) && plusFlag {
		return isEqualLengthRegexMatch(newRegex, input)
	}

	// Characters match
	return handlePlus(regex, input[1:], true)
}

func handleAsterisk(regex, input string) (isMatch bool) {
	if isBaseCaseMet, result := isBaseCase(regex, input); isBaseCaseMet {
		return result == "true"
	}

	newRegex := strings.Replace(regex, regex[0:2], "", 1)
	if isEqualLengthRegexMatch(newRegex, input) {
		return true
	}

	if !isCharacterMatch(regex[0], input[0]) {
		return isEqualLengthRegexMatch(newRegex, input)
	}

	return handleAsterisk(regex, input[1:])
}

func handleQuestionMark(regex, input string) (isMatch bool) {
	if isCharacterMatch(regex[0], input[0]) {
		input = input[1:]
	}

	newRegex := strings.Replace(regex, regex[0:2], "", 1)
	return isEqualLengthRegexMatch(newRegex, input)
}

func isBaseCase(regex, input string) (isMet bool, result string) {
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

// Because of base cases being checked before this function is called, neither argument will ever be ""
func isCharacterMatch(regexChar, inputChar byte) (isMatch bool) {
	return (regexChar == inputChar) || (string(regexChar) == ".")
}

func isRawCharacterMatch(regexChar, inputChar byte) (isMatch bool) {
	return regexChar == inputChar
}

func parseInput() (regex, input string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fullInput := scanner.Text()
	splitInput := strings.Split(fullInput, "|")

	return splitInput[0], splitInput[1]
}
