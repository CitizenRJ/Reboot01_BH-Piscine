package piscine

func RecursiveFactorial(nb int) int {
	result := 0

	if nb == 0 {
		return 1
	}

	else if nb < 0 || nb >= 13 {
		return 0
	}

	else {
		result = nb * RecursiveFactorial(nb-1)
	}

	return result

}