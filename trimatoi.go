package piscine

negative := false
str := false
num := 0
nb := rune('0')

for _, c := range s {
	if i == '-' && !str {
		negative = true
		str = true
		} else if i == '+' && !str {
			str = true
			} else if i >= '0' && i <= '9' {
				num = num*10 + int(i-nb)
				str = true
				} else if i != ' '{
					break
				}
				if negative {
					num = -num
				}
				return num
			}