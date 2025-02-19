package output

import (
	"fmt"
	"strings"
)

func Splitt(content string) [][]string {
	slice1 := []string{}
	result := [][]string{}
	line := ""
	// Splitting the Content with \n
	for _, x := range content {
		if x == '\n' {
			slice1 = append(slice1, line)
			line = ""
		} else if x != '\r' {
			line += string(x)
		}
	}
	if line != "" {
		slice1 = append(slice1, line)
		line = ""
	}

	// Making a slice of slices
	for i := 1; i < len(slice1); i += 8 {
		end := i + 8
		if end > len(slice1) {
			break
		}
		temp := slice1[i:end]
		result = append(result, temp)
		i++
	}

	return result
}

func Split_with_newline(input string) []string {
	var words []string
	word := ""

	for _, x := range input {
		if x != '\n' {
			word += string(x)
		} else {
			if word != "" {
				words = append(words, word)
				word = ""
			} else {
				words = append(words, "\n")
			}
		}
	}
	if word != "" {
		words = append(words, word)
	}
	if input!=""&& input[len(input)-1] == '\n' {
		words = append(words, "\n")
	}
	return words
}

func Find(str, sub_str string) []int {
	if str == sub_str{
		temp := []int{}
		temp = append(temp, 0)
		temp = append(temp, len(str)-1)
		// fmt.Println(temp)
		return temp
	}
	var indexes []int
	start := 0
	var index int

	for {
		index = strings.Index(str[start:], sub_str)
		if index == -1 {
			break
		}
		index += start
		indexes = append(indexes, index)
		start = index + len(sub_str)
	}

	return indexes
}

func Contains(slice []int, target int) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}

func Fill_map(content string) map[rune][]string{
	characters := make(map[rune][]string)      
	temp_slice := Splitt(string(content)) 
	char := 32

	// filling the map
	for _, x := range temp_slice {
		characters[rune(char)] = x
		char++
	}
	return characters
}

func Check_newLines(input string) bool{
	count := 0
	for _, x := range input {
		if x == '\n' {
			count++
		}
	}
	if count == len(input) {
		for count > 0 {
			fmt.Println()
			count--
		}
		return true
	}
	return false
}

func Check_input_subStr(input, sub_str string) bool{
	for _, x := range input {
		if !(x >= ' ' && x <= '~') && x != '\n' {
			fmt.Println("The string contains an unprintable character.")
			return true
		}
	}
	for _, x := range sub_str {
		if !(x >= ' ' && x <= '~'){
			fmt.Println("The substring contains an unprintable character.")
			return true
		}
	}
	return false
}

func Draw(input, user_color, sub_str string, characters map[rune][]string, index []int) []string {
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
	result := make([]string, 8)
	

	colorRanges := make([][2]int, 0)
	for _, idx := range index {
		start := idx
		end := idx + len(sub_str)
		colorRanges = append(colorRanges, [2]int{start, end})
	}
	// fmt.Println(colorRanges)

	for i, x := range input {
		if temp, exist := characters[x]; exist {
			for j := 0; j < 8; j++ {
				isColored := false
				for _, r := range colorRanges {
					if i >= r[0] && i < r[1] { 
						isColored = true
						break
					}
				}
				if isColored {
					result[j] += user_color + temp[j] + colors["reset"]
				} else {
					result[j] += temp[j]
				}
			}
		}
	}

	return result
}
