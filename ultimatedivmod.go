package piscine

func UltimateDivMod(a *int, b *int) {
	if *b == 0 {
		return
	}

	if *a < 0 {
		*a = -*a
	}

	if (*a % *b) == 0 {
		*a /= *b
		*b = *a % *b
	} else {
		return
	}
}
