package functions

import (
	colors "ascii/tools"
)

func PrintAsciiColor(SplitSlice []string, MapAscii map[rune][]string, argument, colorCode string) string {
	var result string

	for _, word := range SplitSlice {
		indexes := IndexAll(word, argument)
		if word != "" {
			for line := 0; line < 8; line++ {
				for i, char := range word {
					if asciiLines, exists := MapAscii[char]; exists {
						asciiLine := asciiLines[line]

						// Vérifier si ce caractère fait partie d'un argument coloré
						shouldColor := false
						for _, indx := range indexes {
							if i >= indx && i < indx+len(argument) {
								shouldColor = true
								break
							}
						}

						if shouldColor {
							result += (colorCode + asciiLine + colors.RESET)
						} else {
							result += (asciiLine)
						}
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
