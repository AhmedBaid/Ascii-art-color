package functions

func PrintAsciiColor(Splitslice,argument []string, Mapascii map[rune][]string, color string) string {
	result := ""
	for _, word := range Splitslice {
		if word != "" {
			for line := 0; line < 8; line++ {
				for i, char := range word {
					if Valeur, exist := Mapascii[char]; exist {
						if i<len(argument[0])&& Splitslice[0][i]==argument[0][i] {
							result += colorMap[color] + (Valeur[line]) + RESET
						}else{
							result += (Valeur[line])
						}
						
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
