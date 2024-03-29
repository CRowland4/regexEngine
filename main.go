package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	StartMessage = "\nWelcome to the Go regex engine! The engine supports the following chcaracters: [^, $, \\, ., ?, *, +]. Press CTRL+C to exit."
)

func main() {
	fmt.Println(StartMessage)
	fmt.Println(Red + "+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n\n" + Reset)
	for {
		regex, input := parseInput()
		if isRegexMatch(regex, input) {
			fmt.Println(Green + "Match!\n" + Reset)
		} else{
			fmt.Println(Red + "Not a match.\n" + Reset)
		}
	}
}

func isRegexMatch(regex, input string) (isMatch bool) {
	if isBaseCaseMet, result := isBaseCase(regex, input); isBaseCaseMet {
		return result == "true"
	}

	if strings.HasPrefix(regex, "^") {
		return isEqualLengthMatch(regex[1:], input)
	}
	if isEqualLengthMatch(regex, input) {
		return true
	}

	return isRegexMatch(regex, input[1:])
}

func isEqualLengthMatch(regex, input string) (isMatch bool) {
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

	return isEqualLengthMatch(regex[1:], input[1:])
}

func handleEscapeCharacter(regex, input string) (isMatch bool) {
	if !isRawCharacterMatch(regex[1], input[0]) {
		return isEqualLengthMatch(regex, input[1:])
	}

	newRegex := strings.Replace(regex, regex[0:2], "", 1)
	return isEqualLengthMatch(newRegex, input[1:])
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
	if plusFlag && isEqualLengthMatch(newRegex, input) {
		return true
	}

	if !isCharacterMatch(regex[0], input[0]) && !plusFlag {
		return false
	}
	if !isCharacterMatch(regex[0], input[0]) && plusFlag {
		return isEqualLengthMatch(newRegex, input)
	}

	// Characters match
	return handlePlus(regex, input[1:], true)
}

func handleAsterisk(regex, input string) (isMatch bool) {
	if isBaseCaseMet, result := isBaseCase(regex, input); isBaseCaseMet {
		return result == "true"
	}

	newRegex := strings.Replace(regex, regex[0:2], "", 1)
	if isEqualLengthMatch(newRegex, input) {
		return true
	}

	if !isCharacterMatch(regex[0], input[0]) {
		return isEqualLengthMatch(newRegex, input)
	}

	return handleAsterisk(regex, input[1:])
}

func handleQuestionMark(regex, input string) (isMatch bool) {
	if isCharacterMatch(regex[0], input[0]) {
		input = input[1:]
	}

	newRegex := strings.Replace(regex, regex[0:2], "", 1)
	return isEqualLengthMatch(newRegex, input)
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

func isCharacterMatch(regexChar, inputChar byte) (isMatch bool) {
	return (regexChar == inputChar) || (string(regexChar) == ".")
}

func isRawCharacterMatch(regexChar, inputChar byte) (isMatch bool) {
	return regexChar == inputChar
}

func parseInput() (regex, input string) {
	scanner := bufio.NewScanner(os.Stdin)
	var fullInput string
	for {
		fmt.Print("Input a regex patter and string in the format '<regex patter>|<string>'\n> ")
		scanner.Scan()
		fullInput = scanner.Text()

		if strings.Count(fullInput, "|") != 1 {
			fmt.Println(Red + "Incorrect input!\n" + Reset)
		}else{ break }

	}

	splitInput := strings.Split(fullInput, "|")

	return splitInput[0], splitInput[1]
}
