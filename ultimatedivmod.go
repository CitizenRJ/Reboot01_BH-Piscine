package piscine

func UltimateDivMod(a *int, b *int) {
	if *b == 0 {
		return
	}

	*a = *a / *b
	*b = *a % *b
}
