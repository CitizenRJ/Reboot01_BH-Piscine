package piscine

func RecursiveFactorial(nb int) int {
	result := 0

	if nb == 0 {
		return 1
	}

	if nb < 0 || nb >= 13 {
		return 0
	}

	{
		result = nb * RecursiveFactorial(nb-1)
	}

	return result

}
