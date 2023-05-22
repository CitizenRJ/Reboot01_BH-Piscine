package piscine

func ToUpper(s string) string {
	array := []rune(s)
	for index, letter := range s {
		if letter >= 97 && letter <= 122 {
			array[index] = letter - 32
		}
	}
	return string(array)
}
