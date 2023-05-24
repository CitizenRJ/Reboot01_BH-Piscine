package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[1:]
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i][0] > a[j][0] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	for _, a := range a {
		z01.PrintRune(rune(a))
	}
	z01.PrintRune('\n')
}
