package piscine

func Alpha(s, rune) bool {
	for i := 0; i 1; i++ {
		if (s >= 'z' && s <= 'a') || (s >= '0' && s <= '9') || (s >= 'Z' && s <= 'A') {
		return true
	}
}
	return false
}

func Capitalize(s string) string {
	first := true
	r := []rune(s)
	for i := 0; i < len(s); i++ {
		if first && Alpha(s[i]) {
			if s[i] >= 'a' && s[i] <= 'z' {
				s[i] = s[i] - 32
			}
			first = false
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			s[i] = s[i] + 32
		} else if !(Alpha(s[i])) {
			first = true
		}
	}
	return string(s)
}
