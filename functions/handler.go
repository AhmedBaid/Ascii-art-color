package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	RESET  = "\033[0m"
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
	PURPLE = "\033[35m"
	CYAN   = "\033[36m"
	WHITE  = "\033[37m"
)

var colorMap = map[string]string{
	"red":    RED,
	"green":  GREEN,
	"yellow": YELLOW,
	"blue":   BLUE,
	"purple": PURPLE,
	"cyan":   CYAN,
	"white":  WHITE,
}

func Handler(argument, banner, fileName, color, something string) {
	fileName = strings.ToLower(fileName)
	file, err := os.Open(fileName)
	color = strings.ToLower(color)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Store ASCII characters from the file
	SlicesAscii := []string{}
	MapAscii := map[rune][]string{}
	count := 0
	espace := ' '
	myscanner := bufio.NewScanner(file)
	for myscanner.Scan() {
		text := myscanner.Text()
		if text != "" {
			SlicesAscii = append(SlicesAscii, text)
			count++
		}
		if count == 8 {
			MapAscii[espace] = SlicesAscii
			espace++
			SlicesAscii = []string{}
			count = 0
		}
	}

	if Isprintable(argument) {
		fmt.Println("Isprintable characters not allowed")
		return
	}

	// Split the input on "\n" to handle multiline ASCII text

	// Handle color highlighting
	if color != "" {
		Splitslice := strings.Split(something, "\\n")
		// Generate ASCII art using PrintAscii function
		var asciiOutput string
		if strings.ReplaceAll(something, "\\n", "") == "" {
			for i := 0; i < strings.Count(something, "\\n"); i++ {
				asciiOutput += "\n"
			}
		} else {
			asciiOutput = PrintAscii(Splitslice, MapAscii)
		}
		if argument != "" && something != "" {
			if colorCode, exists := colorMap[color]; exists {
				coloredAscii := PrintAsciiColor(Splitslice, MapAscii, argument, colorCode)
				fmt.Println(coloredAscii)
			} else {
				fmt.Println("Invalid color specified")
				return
			}
		} else if something != "" && argument == "" {
			asciiOutput = PrintAscii(Splitslice, MapAscii)
			fmt.Println(colorMap[color]+asciiOutput)
		}
	}

	// Save to file if banner is provided
	if banner != "" && argument != "" {
		Splitslice := strings.Split(argument, "\\n")
		var lastResult string
		if strings.ReplaceAll(argument, "\\n", "") == "" {
			for i := 0; i < strings.Count(argument, "\\n"); i++ {
				lastResult += "\n"
			}
		} else {
			lastResult = PrintAscii(Splitslice, MapAscii)
		}
		err := os.WriteFile(banner, []byte(lastResult), 0o644)
		if err != nil {
			fmt.Println("Error", err)
			return
		}
	}
	if argument != "" && banner == "" && color == "" {
		Splitslice := strings.Split(argument, "\\n")
		var lastResult string
		if strings.ReplaceAll(argument, "\\n", "") == "" {
			for i := 0; i < strings.Count(argument, "\\n"); i++ {
				lastResult += "\n"
			}
		} else {
			lastResult = PrintAscii(Splitslice, MapAscii)
		}
		fmt.Println(lastResult)
	}
}
