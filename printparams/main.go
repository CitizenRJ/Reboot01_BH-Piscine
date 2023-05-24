package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[1:]
	for _, i := range a {
		for j := 0; j < len(i); j++ {
			z01.PrintRune(rune(i[j]))
		}
		z01.PrintRune('\n')
	}
}
