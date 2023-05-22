package piscine

func Index(s string, toFind string) int {
	x := []rune(s)
	y := []rune(toFind)
	index := -1
	for i := 0; i < len(x); i++ {
		if len(y)+i <= len(x) {
			if s[i:i+len(y)] == toFind {
				index = i
				break
			}
		}
	}
	return index
}
