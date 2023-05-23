package piscine

func TrimAtoi(s string) int {
	x := 0
	y := 1
	for _, i := range s {
		if i == '-' && x == '0' {
			y = -1
		}
		if i >= '0' && i <= '9' {
			x = x*10 + int(i-48)
		}
	}
	if y {
		return -x
	} else {
		return x
	}
}
