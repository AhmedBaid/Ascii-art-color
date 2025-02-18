package Mosdef

// Color function checks the flag and returns the corresponding RGB color code as an array of three integers.
func Color(flag string) [3]int {
	// Initialize an empty RGB array with default values [0, 0, 0].
	RGB := [3]int{}

	// Check if the flag has a length greater than 8 and starts with "--color=".
	if len(flag) > 8 && flag[0:8] == "--color=" {
		// Check which color is specified after "--color=" and assign corresponding RGB values.
		if flag[8:] == "red" {
			RGB[0] = 255  // Red
			RGB[1] = 0    // Green
			RGB[2] = 0    // Blue
		} else if flag[8:] == "blue" {
			RGB[0] = 0
			RGB[1] = 0
			RGB[2] = 255  // Blue
		} else if flag[8:] == "green" {
			RGB[0] = 0
			RGB[1] = 255  // Green
			RGB[2] = 0
		} else if flag[8:] == "coral" {
			RGB[0] = 255  // Red
			RGB[1] = 127  // Green
			RGB[2] = 80   // Blue
		} else if flag[8:] == "pink" {
			RGB[0] = 255  // Red
			RGB[1] = 0    // Green
			RGB[2] = 127  // Blue
		} else if flag[8:] == "purple" {
			RGB[0] = 76   // Red
			RGB[1] = 0    // Green
			RGB[2] = 153  // Blue
		} else if flag[8:] == "yellow" {
			RGB[0] = 255  // Red
			RGB[1] = 255  // Green
			RGB[2] = 0    // Blue
		} else if flag[8:] == "orange" {
			RGB[0] = 255  // Red
			RGB[1] = 69   // Green
			RGB[2] = 0    // Blue
		} else if flag[8:] == "brown" {
			RGB[0] = 139  // Red
			RGB[1] = 69   // Green
			RGB[2] = 19   // Blue
		} else {
			// If the color is not recognized, assign an invalid RGB color value.
			RGB[0] = 10
			RGB[1] = 10
			RGB[2] = 10
		}
	} else {
		// If the flag does not start with "--color=", assign an invalid RGB color value.
		RGB[0] = 20
		RGB[1] = 20
		RGB[2] = 20
	}
	// Return the RGB array.
	return RGB
}
