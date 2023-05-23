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
	for i = 0; i < len(runes); i++ {
		if runes[i] == ' ' {
			first = true
			continue
		}
		if first {
			if runes[i] >= 'A' && runes[i] <= 'Z' {
				runes[i] = runes[i] - 'A' + 'a'
			}
			first = false
			continue
		}
		if runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] = runes[i] - 'a' + 'A'
		}
	}
	return string(runes)
}
