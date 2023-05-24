package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[1:]
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	for j := 1; j < len(a); j++ {
		for _, k := range a[j] {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')
	}

}
