package Mosdef

// Index function checks if the integer 'j' exists in the slice 'sl'
func Index(j int, sl []int) bool {
	// Loop through each element 'k' in the slice 'sl'
	for _, k := range sl {
		if j == k {
			return true
		}
	}
	return false
}
