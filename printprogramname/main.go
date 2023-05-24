package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i, arg := range os.Args[0] {
		if i > 1 {
			z01.PrintRune(arg)
		}
	}
	z01.PrintRune('\n')
}
