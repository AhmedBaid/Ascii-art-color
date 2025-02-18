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
	// fmt.Println("argument", argument)
	// fmt.Println("banner", banner)
	// fmt.Println("fileName", fileName)
	// fmt.Println("color", color)
	// fmt.Println("something", something)
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

	if len(argument) == 0 {
		return
	}

	if Isprintable(argument) {
		fmt.Println("Isprintable characters not allowed")
		return
	}

	// Split the input on "\n" to handle multiline ASCII text
	Splitslice := strings.Split(something, "\\n")
	var asciiOutput string

	// Generate ASCII art using PrintAscii function
	if strings.ReplaceAll(something, "\\n", "") == "" {
		for i := 0; i < strings.Count(something, "\\n"); i++ {
			asciiOutput += "\n"
		}
	} else {
		asciiOutput = PrintAscii(Splitslice, MapAscii)
	}

	if color != "" {
		if col, exists := colorMap[color]; exists {
			handleColor(argument,something,col,MapAscii)
			return
		} else {
			fmt.Println("Invalid color specified")
			return
		}
	}

	// Save to file or print output
	if banner != "" {
		err := os.WriteFile(banner, []byte(asciiOutput), 0o644)
		if err != nil {
			fmt.Println("Error", err)
			return
		}
	} else {
		fmt.Println(asciiOutput)
	}
}
