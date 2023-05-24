package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[1:]
	for i := len(a); i >= 1; i-- {
		b := os.Args[i]
		for _, j := range b {
			z01.PrintRune(rune(j))
		}
		z01.PrintRune('\n')
	}
}
