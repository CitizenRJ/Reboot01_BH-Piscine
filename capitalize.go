package piscine

func Alpha(s, rune) bool {
	if (s > 'z' || s < 'a') && (s < '0' || s > '9') && (s > 'Z' || s < 'A') {
		return false
	}
	return true
}

func Capitalize(s string) string {
	r := []rune(s)
	first := true
	for i := range runes {
		if check(runes[i])== true && first {
			if runes[i] >= 'a' && runes [i] <= 'z' {
				runes[i] -= 32
			}
			first = false
			}

		else if runes[i] >= 'A' && runes[i] <= 'Z' {
			ruines[i] += 32
			}
			else if check(runes[i]) == false {
				first = true
			}
	}
	return string(runes)
}
