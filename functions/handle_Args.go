package functions

import (
	"fmt"
	"os"
	"strings"
)

func Handle_Args(Args []string) {
	if len(os.Args) < 2 || len(os.Args) > 5 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <something to be colored> \"something\"")
		return
	}
	argument := ""
	fileName := "files/standard.txt"
	banner := ""
	color := ""
	something := ""

	switch {
	case len(os.Args) == 2:
		argument = os.Args[1]

	case len(os.Args) == 3:
		if strings.HasPrefix(os.Args[1], "--color=") {
			color = os.Args[1][8:]
			something = os.Args[2]
			if color == "" {
				fmt.Println("you didn't specify the color")
				return
			}
		} else if strings.HasPrefix(os.Args[1], "--output=") {
			argument = os.Args[2]
			if strings.HasSuffix(os.Args[1], ".txt") {
				banner = os.Args[1][9:]
			} else {
				fmt.Println("Invalid output file name: usage => banner.txt")
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
			something = os.Args[3]
			argument = os.Args[2]
			if color == "" {
				fmt.Println("you didn't specify the color")
				return
			}
		} else if strings.HasPrefix(os.Args[1], "--output=") {
			argument = os.Args[2]
			if strings.HasSuffix(os.Args[1], ".txt") {
				banner = os.Args[1][9:]
			} else {
				fmt.Println("Invalid output file name: usage => banner.txt")
				return
			}
			if strings.HasSuffix(os.Args[3], ".txt") {
				fileName = "files/" + os.Args[3]
			} else {
				fileName = "files/" + os.Args[3] + ".txt"
			}
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <something to be colored> \"something\"")
			return
		}
	case len(os.Args) == 5:
		if strings.HasPrefix(os.Args[1], "--color=") {
			color = os.Args[1][8:]
			argument = os.Args[2]
			something = os.Args[3]
			fileName = os.Args[4]
			if color == "" {
				fmt.Println("you didn't specify the color")
				return
			}
			if strings.HasSuffix(os.Args[4], ".txt") {
				fileName = "files/" + os.Args[4]
			} else {
				fileName = "files/" + os.Args[4] + ".txt"
			}
		} else if strings.HasPrefix(os.Args[1], "--output=") {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER] \n\nEX: go run . --output=<fileName.txt> something standard")
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <something to be colored> \"something\"")
			return
		}
	default:
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <something to be colored> \"something\"")
		return
	}
	Handler(argument, banner, fileName, color, something)
}
