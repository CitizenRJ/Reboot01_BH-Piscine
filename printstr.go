package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {
	byt := []byte(str)
	for _, word := range byt {
		z01.PrintRune(rune(word))
	}
}
