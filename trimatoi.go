package piscine

func TrimAtoi(s string) int {
	x := 1
	y := 0

	for _, i := range s {
		if i >= '0' && i <= '9' {
			j := int(i - '0')
			y = y*10 + j
		} else if i == '-' && y == 0 {
			x = -1
		}
	}
	return x * y
}
