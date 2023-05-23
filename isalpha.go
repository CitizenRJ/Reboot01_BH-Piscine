package piscine

func IsAlpha(s string) bool {
	for _, i := range s {
		if i > 'z' || i < '0' {
			return false
		}
	}
	return true
}
