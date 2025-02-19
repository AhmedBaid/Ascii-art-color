package output

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Data struct {
	FileName   string
	UserColor  string
	Input      string
	SubStr     string
	FileFormat string
	State      bool
	Err        string
}

func Fs(cleanInput []string, characters map[rune][]string) string {
	var final []string
	var found []int
	rest := ""
	for i := 0; i < len(cleanInput); i++ {
		if cleanInput[i] != "\n" {

			final = Draw(cleanInput[i], "", "", characters, found)

			for _, line := range final {
				rest += line + "\n"
			}
		} else {
			rest += "\n"
		}
	}

	return rest
}

func Output(cleanInput []string, characters map[rune][]string, fileName string) string {
	var final []string
	var found []int
	rest := ""
	for i := 0; i < len(cleanInput); i++ {
		if cleanInput[i] != "\n" {

			final = Draw(cleanInput[i], "", "", characters, found)

			for _, line := range final {
				rest += line + "\n"
			}
		} else {
			rest += "\n"
		}
	}
	err := os.WriteFile(fileName, []byte(rest), 0o644)
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println()
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
	}
	return rest
}

func Color(cleanInput []string, userColor string, cleanSubStr []string, characters map[rune][]string, matchFound bool) {
	var final []string
	j := 0
	colors := map[string]string{
		"default": "\033[39m",
		"black":   "\033[30m",
		"gray":    "\033[37m",
		"red":     "\033[91m",
		"green":   "\033[92m",
		"orange":  "\033[38;5;208m",
		"blue":    "\033[94m",
		"white":   "\033[97m",
		"reset":   "\033[0m",
		"yellow":  "\033[33m",
	}
	
	if colorCode, exists := colors[userColor]; exists {
		for i := 0; i < len(cleanInput); i++ {
			found := Find(cleanInput[i], cleanSubStr[j])
			if reflect.DeepEqual(cleanInput, cleanSubStr) {
				found = Find(cleanInput[i], cleanInput[i])
			}
			
			if cleanInput[i] != "\n" {
				if matchFound {
					if reflect.DeepEqual(cleanInput, cleanSubStr) {
						final = Draw(cleanInput[i], colorCode, cleanInput[i], characters, found)
					} else {
						final = Draw(cleanInput[i], colorCode, cleanSubStr[j], characters, found)
					}

					if j < len(cleanSubStr)-1 {
						j++
					}
				} else {
					final = Draw(cleanInput[i], colorCode, "", characters, found)
				}
				for _, line := range final {
					fmt.Println(line)
				}
			} else {
				fmt.Print(cleanInput[i])
			}
		}
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
	}
}

func CleanOutput(args []string) Data {
	fileName := ""
	colorFlag := ""
	userColor := ""
	input := ""
	subStr := ""
	state := false
	err := ""
	fileFormat := "standard.txt"
	if len(args) == 1 {
		err = "Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard"
		// Data.err = true
		return Data{fileName, userColor, input, subStr, fileFormat, state, err}
	}
	if strings.HasPrefix(args[1], "--color=") {
		if len(args) < 3 || len(args) > 5 {
			err = "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\""
			// Data.err = true
			return Data{fileName, userColor, input, subStr, fileFormat, state, err}
		}
		if strings.HasPrefix(args[1], "--color=") && strings.HasPrefix(args[2], "--output=") {
			err = "Use one flag"
			// Data.err = true
			return Data{fileName, userColor, input, subStr, fileFormat, state, err}
		}
		if len(args) == 5 {
			colorFlag = args[1]
			subStr = args[2]
			input = args[3]
			if strings.HasSuffix(args[4], ".txt") {
				fileFormat = args[4]
			} else {
				fileFormat = args[4] + ".txt"
			}
		} else if len(args) == 4 {
			colorFlag = args[1]
			subStr = args[2]
			input = args[3]

		} else {
			colorFlag = args[1]
			state = true
			// subStr = args[2]
			input = args[2]
		}
	} else if strings.HasPrefix(args[1], "--output=") {
		if len(args) < 3 || len(args) > 4 {
			err = "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard"
			return Data{}
		}
		if len(args) == 4 {
			fileName = args[1]

			input = args[2]
			if strings.HasSuffix(args[3], ".txt") {
				fileFormat = args[3]
			} else {
				fileFormat = args[3] + ".txt"
			}
		} else if len(args) == 3 {
			fileName = args[1]

			input = args[2]

		}
	} else {
		if len(args) < 2 || len(args) > 3 {
			err = "Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard"
			// Data.err = true
			return Data{fileName, userColor, input, subStr, fileFormat, state, err}
		}
		if len(args) == 3 {
			if strings.HasSuffix(args[2], ".txt") {
				fileFormat = args[2]
			} else {
				fileFormat = args[2] + ".txt"
			}

			input = args[1]
		} else {
			input = args[1]
		}
	}

	if strings.HasPrefix(colorFlag, "--color=") {
		userColor = colorFlag[8:]
	}

	if strings.HasPrefix(fileName, "--output=") {
		fileName = fileName[9:]
	}

	return Data{fileName, userColor, input, subStr, fileFormat, state, err}
}
