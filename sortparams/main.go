package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := len(os.Args) - 1
	for i := 1; i <= a; i++ {
		for j := 1; j <= a; j++ {
			if os.Args[i] < os.Args[j] {
				os.Args[i], os.Args[j] = os.Args[j], os.Args[i]
			}
		}
	}
	for i := 1; i <= a; i++ {
		for _, x := range os.Args[i] {
			z01.PrintRune(x)
		}
		z01.PrintRune('\n')
	}
}
