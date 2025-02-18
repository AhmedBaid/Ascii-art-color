package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	mosdef "Mosdef/funcs"
)

func main() {
	var argument string
	var subString string
	RGB := [3]int{200, 200, 200} // Default RGB color (white)

	// Call CheckArgs to process command-line arguments and assign values to RGB, argument, and subString
	RGB, argument, subString = mosdef.CheckArgs(os.Args, argument, RGB, subString)

	// Reading the ASCII art file "standard.txt"
	ascii, er1 := os.ReadFile("standard.txt")
	// Check if the file size is as expected, otherwise print an error message
	if len(ascii) != 6623 {
		fmt.Println("data error please check the standard.txt file") // Print error if the file size is incorrect
		return
	}
	// If an error occurs while reading the file, log the error and terminate the program
	if er1 != nil {
		log.Fatal(er1)
	}

	// Split the content of the file into separate lines by using newline "\n" as a delimiter
	lines := strings.Split(string(ascii), "\n")
	// Split the input argument (the sentence) into lines based on the literal "\n" for ASCII art processing
	sentence := strings.Split(argument, "\\n")

	// Call GenerateAscii from the mosdef package to generate the ASCII art with the processed data
	mosdef.GenerateAscii(lines, sentence, subString, RGB)
}
