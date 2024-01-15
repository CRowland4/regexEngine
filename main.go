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

func regexMatch(regex, input string) (isMatch bool) {
	// Regex starts with ^ and ends with $: input must match inner text of regex exactly
	if strings.HasPrefix(regex, "^") && strings.HasSuffix(regex, "$") {
		return regex[1:len(regex)-1] == input // TODO this doesn't account for wildcards.
	}

	// Empty regex, or regex that has had everything consumed but ^: always counts as a match
	if (regex == "") || (regex == "^") {
		return true
	}

	// Regex ends with $, doesn't start with ^, but the input has more characters: trim input down
	if strings.HasSuffix(regex, "$") && (len(regex)-1 < len(input)) {
		return regexMatch(regex, input[1:])
	}

	// No ^ in regex, and regex has been consumed to the dollar sign: match if input has been fully consumed
	if regex == "$" {
		return input == ""
	}

	// Never a match if the input is blank and regex still has characters remaining that aren't ^ and $
	if input == "" {
		return false
	}

	// Regex starts with ^, but the next character doesn't match the next character of input
	if strings.HasPrefix(regex, "^") && !doCharactersMatch(regex[1], input[0]) {
		return false
	}

	// Regex starts with ^ and next character matches the next character of input
	if strings.HasPrefix(regex, "^") && doCharactersMatch(regex[1], input[0]) {
		regex = strings.Replace(regex, string(regex[1]), "", 1)
		return regexMatch(regex, input[1:])
	}

	if doCharactersMatch(regex[0], input[0]) {
		return regexMatch(regex[1:], input[1:])
	}

	return regexMatch(regex, input[1:])
}

func doCharactersMatch(regexChar, inputChar byte) (isMatch bool) {
	return (string(regexChar) == ".") || (regexChar == inputChar)
}

func parseInput() (regex, input string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fullInput := scanner.Text()
	splitInput := strings.Split(fullInput, "|")

	regex, input = splitInput[0], splitInput[1]
	return splitInput[0], splitInput[1]
}
