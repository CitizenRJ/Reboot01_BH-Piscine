package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	i := os.Args[1:]
	for _, j := range i {
		for k := 0; k < len(j); k++ {
			z01.PrintRune(rune(i[k]))
		}
		z01.PrintRune('\n')
	}
}
