package functions

func PrintAscii(Splitslice []string, Mapascii map[rune][]string) string {
	result := ""
	for _, word := range Splitslice {
		if word != "" {
			for line := 0; line < 8; line++ {
				for _, char := range word {
					if Valeur, exist := Mapascii[char]; exist {
						result += Valeur[line]
					}
				}
				result += "\n"
			}
		} else {
			result += "\n"
		}
	}
	return result
}
