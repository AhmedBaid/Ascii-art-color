package functions

import (
	"strings"
)

func PrintAsciiColor(Splitslice []string, MapAscii map[rune][]string, argument, colorCode string) string {
	result := ""

	for _, word := range Splitslice {
		index := strings.Index(word, argument)
		if word != "" {
			for line := 0; line < 8; line++ {
				for i, char := range word {
					if asciiLines, exists := MapAscii[char]; exists {
						asciiLine := asciiLines[line]

						if strings.Contains(word, argument) && index != -1 && i >= index && i < index+len(argument) {
							asciiLine = colorCode + asciiLine + RESET
						} else {
							asciiLine = asciiLine + RESET
						}
						result += asciiLine
					}
				}
				result += "\n"
			}
		} else {
			result += "\n"
		}
	}
	return result
}
