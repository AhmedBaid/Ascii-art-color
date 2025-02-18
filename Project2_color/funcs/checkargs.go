package Mosdef

// CheckArgs function processes command-line arguments and sets the RGB color, main argument, and substring.
// It returns the RGB array, the main argument, and the substring.
func CheckArgs(Args []string, argument string, RGB [3]int, subString string) ([3]int, string, string) {
	var color string // Variable to hold the color flag

	// Check the length of command-line arguments
	if len(Args) == 4 { // Case when 4 arguments are provided
		argument = Args[3] // Set the main argument from the last argument
		color = Args[1]    // Set the color from the second argument
		subString = Args[2] // Set the substring from the third argument
		RGB = Color(color)  // Get the RGB values based on the color flag
	} else if len(Args) == 3 { // Case when 3 arguments are provided
		argument = Args[2] // Set the main argument from the last argument
		color = Args[1]    // Set the color from the second argument
		RGB = Color(color) // Get the RGB values based on the color flag
	} else if len(Args) == 2 { // Case when 2 arguments are provided
		argument = Args[1] // Set the main argument from the last argument
		if argument[0:8] == "--color=" { // Check if the argument is a color flag
			argument = "ù" // Set argument to a placeholder if it's a color flag
		}
	}

	// Return the RGB array, the main argument, and the substring
	return RGB, argument, subString
}
