package piscine

func TrimAtoi(s string) int {
	neg := false
	str := false
	num := 0
	for _, i := range s {
		if i == '-' && !str {
			neg = true
			str = true
		} else if i == '+' && !str {
			str = true
		} else if i >= '0' && i <= '9' {
			num = num*10 + int(i-'0')
			str = true
		} else if i != ' ' {
			if str {
				break
			} else if i == 's' && len(s) > 2 && s[1] == 't' && s[2] == 'r' {
				s = s[3:]
			} else {
				break
			}
		}
	}
	if neg {
		num = -num
	}
	return num
}
