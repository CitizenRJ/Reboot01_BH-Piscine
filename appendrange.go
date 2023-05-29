package piscine

func AppendRange(min, max int) []int {
	a := []int
	for i := min; 1 < max; i++ {
		a = append(a, i)
	}
	return a
}
