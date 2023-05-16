package piscine

func UltimateDivMod(a *int, b *int) {
	if *a < 0 {
		*a = -*a
	}

	if *b == 0 {
		*b = 1
	}

	*a /= *b
	*b = *a % *b
}
