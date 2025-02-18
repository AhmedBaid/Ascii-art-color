package functions

import (
	"fmt"
	"strings"
)

func handleColor(argument, something, color string, Mapascii map[rune][]string) {
	Splitslice := strings.Split(something, "\\n")
	asciiOutput := PrintAscii(Splitslice, Mapascii)

	argumentLower := strings.ToLower(argument)
	somethingLower := strings.ToLower(something)

	asciiLines := strings.Split(asciiOutput, "\n")
	coloredLines := make([]string, len(asciiLines))

	startIndex := strings.Index(somethingLower, argumentLower)

	if startIndex != -1 {
		argRunes := []rune(something)[startIndex : startIndex+len(argument)]
		argString := string(argRunes)

		for i, line := range asciiLines {
			coloredLines[i] = strings.Replace(line, argString, RESET+argString+color, 1)
		}
	} else {
		coloredLines = asciiLines
	}

	finalOutput := strings.Join(coloredLines, "\n")
	fmt.Println(finalOutput)
}
