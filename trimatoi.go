package piscine

func TrimAtoi(s string) int {
	neg := false
	str := false
	num := 0
	nb := rune('0')
	for _, i := range s {
		if i == '-' && !str {
			neg = true
			str = true
		} else if i == '+' && !str {
			str = true
		} else if i >= '0' && i <= '9' {
			num = num*10 + int(i-nb)
			str = true
		} else if i != ' ' {
			break
		}
		if neg {
			num = -num
		}
		return num
	}
	return num
}
