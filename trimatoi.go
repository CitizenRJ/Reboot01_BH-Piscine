package piscine

func TrimAtoi(s string) int {
	neg := false
	str := false
	num := 0
	
	for i:= 0, i < len(s); i++ {
		j := s[i]
		if j == '-' && !str {
			neg = true
			str = true
		} else if j == '+' && !str {
			str = true
		} else if j >= '0' && i <= '9' {
			num = num*10 + int(j-'0')
			str = true
		} else if j != ' ' {
			return 0
			
		}
	}
	if neg {
		num = -num
	}
	return num
}
