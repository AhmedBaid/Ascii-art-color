package functions

import (
	"fmt"
	"strings"
)

func handleColor(argument, something, color string, Mapascii map[rune][]string) {
	fmt.Println("color", color)
	var coloredOutput string
	var final string
	if argument != "" && something != "" {
		index := strings.Index(something, argument)
		if index != -1 {
			piecee_colred := color + something[index:index+len(argument)]
			coloredOutput = something[:index] + piecee_colred + something[index+len(argument):]
			final=PrintAsciiColor(piecee_colred,coloredOutput)
		}
	} else if something != "" {
		fmt.Println(color + something + RESET)
		final=PrintAsciiColor()
		return
	} else {
		final=PrintAsciiColor()
	}

	fmt.Println(final)
}
