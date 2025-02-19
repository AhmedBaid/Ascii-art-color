package functions

import (
	"fmt"
	"strings"
)

var (
	Index     int
	ArgString string
)

func PrintAsciiColor(Splitslice []string, MapAscii map[rune][]string, Splitargument []string, argument, colorCode string) string {
	/* fmt.Println("",Splitargument)
	fmt.Println("",Splitslice)
	fmt.Println("",len(Splitslice))
	fmt.Println("",len(Splitargument)) */
	result := ""
	/* fmt.Println("substring",Splitargument)
	   fmt.Println("somthings ",Splitslice) */
	// Join Splitargument and Splitslice to find index
	something := strings.Join(Splitslice, "")
	ArgString = strings.Join(Splitargument, "")
	fmt.Println("substring", something)
	fmt.Println("somthings ", ArgString)

	Index = strings.Index(something, ArgString) //+len(ArgString)-1 // Find position of argument
	fmt.Println(Splitslice)

	for _, word := range Splitslice {
		if word != "" { // Normal ASCII printing
			for line := 0; line < 8; line++ {
				for _, char := range word {
					if asciiLines, exists := MapAscii[char]; exists {
						asciiLine := asciiLines[line]

						// Color only the argument if found
						/* 	if index != -1 && i >= index && i < index+len(argument)-1 {
							fmt.Println("iiii",  i )
							fmt.Println("j",  len(word) )

							fmt.Println("index",  index+len(argument)-1  )

							asciiLine = colorCode + asciiLine + RESET
						} else  { // If argument not found, color everything
							asciiLine = asciiLine + RESET
						} */
						result += asciiLine
					}
				}
				result += "\n"
			}
		} else {
			// If there was an empty line (`\n` in `something`), add a newline
			result += "\n"
		}
	}
	return result
}
