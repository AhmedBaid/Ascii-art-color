package functions

func Isprintable(argument string) bool {
	for _, char := range argument {
		if char < 32 || char > 126 {
			return true
		}
	}
	return false
}
