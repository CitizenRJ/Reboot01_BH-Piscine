package piscine

func Alpha(s, rune) bool {
	if (s > 'z' || s < 'a') && (s < '0' || s > '9') && (s > 'Z' || s < 'A') {
		return false
	}
	return true
}

func Capitalize(s string) string {
	first := true
	r := []rune(s)
	for i := 0; i < len(s); i++ {
		if first && Alpha(rune[i]) {
			if rune[i] >= 'a' && rune[i] <= 'z' {
				rune[i] = rune[i] + 'A' - 'a'
			}
			first = false
		} else if rune[i] >= 'A' && rune[i] <= 'Z' {
			rune[i] = rune[i] + 'a' - 'A'
		} else if !Alpha(rune[i]) {
			first = true
		}
	}
	return string(r)
}
