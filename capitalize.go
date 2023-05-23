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
		if first && Alpha(r[i]) {
			if r[i] >= 'a' && r[i] <= 'z' {
				r[i] = r[i] + 'A' - 'a'
			}
			first = false
		} else if r[i] >= 'A' && r[i] <= 'Z' {
			r[i] = r[i] + 'a' - 'A'
		} else if !Alpha(r[i]) {
			first = true
		}
	}
	return string(r)
}
