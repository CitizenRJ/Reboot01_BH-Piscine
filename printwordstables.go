package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, i := range a {
		for _, a := range i {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}
