package piscine

func isAlphanumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func Capitalize(s string) string {
	runes := []rune(s)
	first := true

	for i := 0; i < len(runes); i++ {
		if isAlphanumeric(runes[i]) {
			if first && runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = runes[i] - ('a' - 'A')
			} else if !first && runes[i] >= 'A' && runes[i] <= 'Z' {
				runes[i] = runes[i] + ('a' - 'A')
			}
			first = false
		} else {
			first = true
		}
	}
	return string(runes)
}
