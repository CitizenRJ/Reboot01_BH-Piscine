package piscine

func TrimAtoi(s string) int {
	x := 1
	y := 0
	z := false

	for _, i := range s {
		if i == '-' && !x {
			x = -1
			z = true
		} else if i == '+' && !z {
			x = true
		} else if i >= '0' && i <= '9' {
			y *= 10
			y += int(i - '0')
			z = true
		} else if i != ' ' {
			break
		}
	}
	return x * y
}
