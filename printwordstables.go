package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, i := range a {
		for _, s := range i {
			z01.PrintRune(s)
		}
		z01.PrintRune('\n')
	}
}
