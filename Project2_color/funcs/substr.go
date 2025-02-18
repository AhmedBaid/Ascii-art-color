package Mosdef

// SubStr function searches for the occurrences of a substring in a string and returns their indices
func SubStr(str, subString string) []int {
	sl := []int{}
	
	// Iterate over each index 'x' in the string 'str'
	for x := range str {
		// Check if the remaining part of 'str' can fit the 'subString' and if the current slice equals the 'subString'
		if x <= len(str)-len(subString) && str[x:x+len(subString)] == subString {
			sl = append(sl, x)
			// Add each subsequent index for the length of the substring to the slice 'sl'
			for h := 0; h < len(subString); h++ {
				sl = append(sl, x+h)
			}
		}
	}
	return sl
}
