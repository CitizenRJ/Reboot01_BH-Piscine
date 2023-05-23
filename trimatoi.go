package piscine

func TrimAtoi(s string) int {
	x := true
	y := false
	z := 0
	var nb rune = 48
	for _, i := range s {
		if i == '-' && x {
			y = true
		}
		if i >= nb && i <= 57 {
			z *= 10
			z += int(i - 48)
			x = false
		}
	}
	if y == true {
		return -z
	} else {
		return z
	}
}
