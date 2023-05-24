package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var a int
	for i := range os.Args {
		a = 1
	}
	for i := 1; i <= a; i++ {
		for j := 1; j <= a; j++ {
			if os.Args[i] < os.Args[j] {
				os.Args[i], os.Args[j] = os.Args[j], os.Args[i]
			}
		}
	}
	for i := range os.Args {
		if i > 0 {
			for _, x := range os.Args[i] {
				z01.PrintRune(x)
			}
			z01.PrintRune('\n')
		}
	}
}
