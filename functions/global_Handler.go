package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	colors "ascii/tools"
)

func Handler(argument, banner, fileName, color, something string) {
	banners := []string{"files/standard.txt", "files/shadow.txt", "files/thinkertoy.txt", "./files/standard.txt", "./files/shadow.txt", "./files/thinkertoy.txt"}
	for _, bnr := range banners {
		if banner == bnr {
			fmt.Println("writing in our files is forbidden")
			return
		}
	}
	fileName = strings.ToLower(fileName)
	color = strings.ToLower(color)
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
			if strings.Contains(argument, `\n`) ||
				strings.Contains(argument, `\t`) ||
				strings.Contains(argument, `\r`) ||
				strings.Contains(argument, `\v`) ||
				strings.Contains(argument, `\f`) {
				fmt.Println("Non-printable sequences (e.g., \\n, \\t) are not considered for coloring")
				return
			}
			if Isprintable(something) {
				fmt.Println("Isprintable characters not allowed")
				return
			}
			if colorCode, exists := colors.ColorMap[color]; exists {
				coloredAscii := PrintAsciiColor(Splitslice, MapAscii, argument, colorCode)
				fmt.Print(coloredAscii)
			} else {
				fmt.Println("Invalid color specified")
				return
			}
		} else if something != "" && argument == "" {
			if Isprintable(something) {
				fmt.Println("Isprintable characters not allowed")
				return
			}
			if colorCode, exists := colors.ColorMap[color]; exists {
				asciiOutput = PrintAscii(Splitslice, MapAscii)
				fmt.Print(colorCode + asciiOutput + colors.RESET)
			} else {
				fmt.Println("Invalid color specified")
				return
			}
		}
	}
	// Save to file if banner is provided
	if banner != "" && argument != "" {
		if Isprintable(argument) {
			fmt.Println("Isprintable characters not allowed")
			return
		}
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
		if Isprintable(argument) {
			fmt.Println("Isprintable characters not allowed")
			return
		}
		Splitslice := strings.Split(argument, "\\n")
		var lastResult string
		if strings.ReplaceAll(argument, "\\n", "") == "" {
			for i := 0; i < strings.Count(argument, "\\n"); i++ {
				lastResult += "\n"
			}
		} else {
			lastResult = PrintAscii(Splitslice, MapAscii)
		}
		fmt.Print(lastResult)
	}
}
