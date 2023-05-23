package piscine

func TrimAtoi(s string) int {
	x := ""
	for _, i := range s {
		if (i >= '0' && i <= '9') || i == '-' {
			if len(x) == 0 && i == '-' {
				x + string(i)
				continue
			}
			if i != '-' {
				x + string(i)
			}
		}
	}
	return Atoi(x)
}
