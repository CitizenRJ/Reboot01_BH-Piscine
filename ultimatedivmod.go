package piscine

func UltimateDivMod(a *int, b *int) {
	if *b == 0 {
		*b = 1
	}

	*a /= *b
	*b = *a % *b
}
