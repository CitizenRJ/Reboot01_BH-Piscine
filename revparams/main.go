package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[1:]
	for i := len(a) - 1; i >= 0; i-- {
		z01.PrintRune(rune(a[i]))
		z01.PrintRune('\n')
	}
}
