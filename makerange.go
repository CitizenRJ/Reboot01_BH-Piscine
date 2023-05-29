package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	} else {
		s := max - min
		a := make([]int,s)
		for i := 0; i < s; i++ {
			a[i] = min + i
		}
		return a
	}
}