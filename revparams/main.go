package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[1:]
	for _, i := len(a) - 1; i >= 0; i-- {
		for j := len(a[i]) - 1; j >= 0; j-- {
			z01.PrintRune(rune(a[i][j]))
		}
		z01.PrintRune('\n')
	}
}
