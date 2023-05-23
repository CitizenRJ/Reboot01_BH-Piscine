package piscine

func Capitalize(s string) string {
	r := []rune(s)
	for index, letter := range r {
		if Alpha(letter) {
			if index == 0 || Alpha(r[index-1]) == false {
			} else if letter >= 'a' && letter <= 'z' {
				r[index] = letter + 32
			}
		}
	}
	return string(r)
}
