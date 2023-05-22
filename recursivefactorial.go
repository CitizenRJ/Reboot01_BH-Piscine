package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		result := 1
		for i := 1; i <= nb; i++ {
			result *= i
		}
		return result
	}
}
