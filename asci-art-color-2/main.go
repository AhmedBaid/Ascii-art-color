package main

import (
	"fmt"
	"os"
	"strings"

	"output/output"
)

func main() {
	args := os.Args
	data := output.CleanOutput(args)
	if data.Err != "" {
		fmt.Println("Error:", data.Err)
		return
	}
	// ARGS
	input := data.Input
	sub_str := data.SubStr
	file_format := strings.ToLower(data.FileFormat)
	userColor := data.UserColor
	fileName := data.FileName
	fmt.Println(fileName)
	state := data.State
	if fileName != "" {
		if strings.HasPrefix(fileName, "standard") || strings.HasPrefix(fileName, "shadow") || strings.HasPrefix(fileName, "thinkertoy") || !strings.HasSuffix(fileName, ".txt") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("")
			fmt.Println("EX: go run . --output=<fileName.txt> something standard")
			return
		}
	}

	input = strings.ReplaceAll(input, "\\n", "\n")
	sub_str = strings.ReplaceAll(sub_str, "\\n", "\n")

	status := output.Check_newLines(input)
	if status {
		return
	}

	status = output.Check_input_subStr(input, sub_str)
	if status {
		return
	}

	content, err1 := os.ReadFile(file_format)
	if err1 != nil {
		fmt.Println("File reading error, ", err1)
		return
	}

	//--------------------------------------------------------------------------------

	characters := output.Fill_map(string(content))
	matchFound := strings.Contains(input, sub_str)
	if state {
		matchFound = true
	}

	cleanInput := output.Split_with_newline(input)
	cleanSubStr := output.Split_with_newline(sub_str)

	if strings.HasPrefix(os.Args[1], "--color=") || userColor != "" {
		if state {
			output.Color(cleanInput, userColor, cleanInput, characters, matchFound)
		} else {
			output.Color(cleanInput, userColor, cleanSubStr, characters, matchFound)
		}
	} else if strings.HasPrefix(os.Args[1], "--output=") || fileName != "" {
		output.Output(cleanInput, characters, fileName)
	} else {
		fmt.Print(output.Fs(cleanInput, characters))
	}
}
