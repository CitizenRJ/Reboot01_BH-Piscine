package piscine

func ToLower(s string) string {
	array := []rune(s)
	for index, letter := range s {
		if letter >= 65 && letter <= 90 {
			array[index] = letter + 32
		}
	}
	return string(array)
}
