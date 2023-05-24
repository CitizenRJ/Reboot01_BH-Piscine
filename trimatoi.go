package piscine

func TrimAtoi(s string) int {
	neg := false
	str := false
	num := 0
	
	if j:= s[0] {
	if j == '-'|| j == '+' || (j >= '0' && j <= '9') {
		num = int(j - '0')
		str = true
	}
}
	for _, j:= range s[1:] {
		if j >= '0' && i <= '9' {
			num = num*10 + int(j-'0')
			str = true
		} else if str {
			break
		} else if j == '-' {
			neg = true
			str = true
		} else if j == '+' {
		str = true
	}
}
	if neg {
		num = -num
	}
	return num
}
