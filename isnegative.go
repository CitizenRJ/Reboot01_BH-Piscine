package main

import "piscine"

func IsNegative(nb int) {
	if nb < 0 {
		piscine.PrintRune('T')
	} else {
		piscine.PrintRune('F')
		}	

func main() {
	piscine.IsNegative(1)
	piscine.IsNegative(0)
	piscine.IsNegative(-1)
}
}