package functions

func IndexAll(str, sub string) []int {
	var indexes []int
	if len(sub) == 0 {
		return indexes
	}

	for i := 0; i <= len(str)-len(sub); i++ {
		if str[i:i+len(sub)] == sub {
			indexes = append(indexes, i)
		}
	}
	return indexes
}
