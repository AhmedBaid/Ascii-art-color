package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/functions"
)

func main() {
	argument := ""
	fileName := "files/standard.txt"
	banner := ""
	color := ""
	switch {
	case len(os.Args) == 2:
		argument = os.Args[1]

	case len(os.Args) == 3:
		if strings.HasPrefix(os.Args[1], "--color=") {
			color = os.Args[1][8:]
			argument = os.Args[2]
		} else if strings.HasPrefix(os.Args[1], "--output=") {
			argument = os.Args[2]
			if strings.HasSuffix(os.Args[1], ".txt") {
				banner = os.Args[1][9:]
			} else {
				fmt.Println("Invalid output file name  : usage => banner.txt")
				return
			}
		} else {
			argument = os.Args[1]
			if strings.HasSuffix(os.Args[2], ".txt") {
				fileName = "files/" + os.Args[2]
			} else {
				fileName = "files/" + os.Args[2] + ".txt"
			}
		}
	case len(os.Args) == 4:
		if strings.HasPrefix(os.Args[1], "--color=") {
			color = os.Args[1][8:]
			argument = os.Args[2]
			if strings.HasSuffix(os.Args[3], ".txt") {
				fileName = "files/" + os.Args[3]
			} else {
				fileName = "files/" + os.Args[3] + ".txt"
			}
		} else if strings.HasPrefix(os.Args[1], "--output=") {
			argument = os.Args[2]
			if strings.HasSuffix(os.Args[1], ".txt") {
				banner = os.Args[1][9:]
			} else {
				fmt.Println("Invalid output file name  : usage => banner.txt")
				return
			}
			if strings.HasSuffix(os.Args[3], ".txt") {
				fileName = "files/" + os.Args[3]
			} else {
				fileName = "files/" + os.Args[3] + ".txt"
			}
		}
	case len(os.Args) == 5:
		if strings.HasPrefix(os.Args[1], "--output=") && strings.HasPrefix(os.Args[2], "--color=") {
			banner = os.Args[1][9:]
			color = os.Args[2][8:]
			argument = os.Args[3]
			if strings.HasSuffix(os.Args[4], ".txt") {
				fileName = "files/" + os.Args[4]
			} else {
				fileName = "files/" + os.Args[4] + ".txt"
			}
		}
	default:
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> --color=<color> something standard")
		return
	}
	functions.Handler(argument, banner, fileName, color)
}
