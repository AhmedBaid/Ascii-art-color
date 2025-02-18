package Mosdef

import (
	"fmt"
)

// GenerateAscii generates and prints colored ASCII art for the provided sentence.
func GenerateAscii(lines, sentence []string, subString string, RGB [3]int) {
	R := RGB[0]
	G := RGB[1]
	B := RGB[2]
	if R == 10 && G == 10 && B == 10 {
		fmt.Println("invalid color")
		return
	} else if R == 20 && G == 20 && B == 20 {
		fmt.Println("invalid syntax")
		return
	}
	// Loop through each line of the sentence
	for r := range sentence {
		start := SubStr(sentence[r], subString)
		x := []int{}
		for _, i := range sentence[r] {
			x = append(x, int(i))
		}
		// Validate the input characters (only allow printable ASCII)
		for ran := range x {
			if x[ran] < 32 || x[ran] > 126 {
				fmt.Println("invalid input")
				return
			}
		}
		// If there's a substring, highlight it with RGB color
		if subString != "" {
			for k := 1; k <= 8; k++ {
				for j, i := range x {
					firstLine := (i-32)*9 + k
					if !Index(j, start) {
						fmt.Print(lines[firstLine])
						if j == len(x)-1 {
							fmt.Print("\n")
						}
					} else if Index(j, start) {
						fmt.Print(fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", R, G, B, lines[firstLine]))
						if j == len(x)-1 {
							fmt.Print("\n")
						}
					} else if subString == "" && (R != 0 || G != 0 || B != 0) {
						fmt.Print(fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", R, G, B, lines[firstLine]))
						if j == len(x)-1 {
							fmt.Print("\n")
						}
					}
				}
			}
		// If no substring, color the entire sentence with the provided RGB	
		} else if subString == "" {
			for k := 1; k <= 8; k++ {
				for j, i := range x {
					firstLine := (i-32)*9 + k
					fmt.Print(fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", R, G, B, lines[firstLine]))
					if j == len(x)-1 {
						fmt.Print("\n")
					}

				}
			}
		}

		// Handle blank lines for readability
		if len(x) == 0 {
			if r != 0 && len(sentence[r-1]) != 0 {
				fmt.Println("")
			} else {
				if r != len(sentence)-1 {
					fmt.Println("")
				}
			}
		}
	}
}
