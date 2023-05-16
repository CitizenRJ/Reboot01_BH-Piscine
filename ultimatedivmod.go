package piscine

import "fmt"

func UltimateDivMod(a *int, b *int) {
	if *b == 0 {
		fmt.Println("Cannot divide by zero")
		return
	}

	*a /= *b
	*b %= *a
}
