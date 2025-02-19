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
	Splitslice := strings.Split(something, "\\n")
	Splitargument := strings.Split(argument, "\\n")
	// Generate ASCII art using PrintAscii function
	var asciiOutput string
	if strings.ReplaceAll(something, "\\n", "") == "" {
		for i := 0; i < strings.Count(something, "\\n"); i++ {
			asciiOutput += "\n"
		}
	} else {
		asciiOutput = PrintAscii(Splitslice, MapAscii)
	}

	// Handle color highlighting
	if color != "" {
		if colorCode, exists := colorMap[color]; exists {
			coloredAscii := PrintAsciiColor(Splitslice, MapAscii, Splitargument, argument, colorCode)

			for i := 0; i < len(coloredAscii); i++ {
				if i*4*8 >= Index+len(ArgString) {
					fmt.Print(colorMap[color] + string(coloredAscii[i]))
				} else {
					fmt.Print(string(coloredAscii[i]))
				}
			}

			// fmt.Println(coloredAscii)
		} else {
			fmt.Println("Invalid color specified")
			return
		}
	} else {
		fmt.Println(asciiOutput)
	}

	// Save to file if banner is provided
	if banner != "" {
		err := os.WriteFile(banner, []byte(asciiOutput), 0o644)
		if err != nil {
			fmt.Println("Error", err)
			return
		}
	}
}
