package piscine

func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func Capitalize(s string) string {
	runes := []rune(s)
	first := false

	for i := 0; i < len(runes); i++ {
		if isAlphanumeric(runes[i]) {
			if first || (i == 0 && runes[i] >= 'a' && runes[i] <= 'z') {
				runes[i] = runes[i] - 32
				first = false
			} else if runes[i] >= 'A' && runes[i] <= 'Z' {
				runes[i] = runes[i] + 32
			}
			first = false
		} else {
			first = true
		}
	}
	return string(runes)
}
