package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	result := 1

	if power%2 == 1 {
		result *= nb
	}

	nb *= nb
	power /= 2

	return RecursivePower(nb, power) * result
}
