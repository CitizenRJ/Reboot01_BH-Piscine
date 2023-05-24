package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[:1]
	for _, i := range a {
		for j := len(i); j > 0; j-- {
			z01.PrintRune(rune(i[j]))
		}
		z01.PrintRune('\n')
	}
}
