package piscine

import "strconv"

func TrimAtoi(s string) int {
	var s string
	for _, i := range s {
		if (i >= '0' && i <= '9') || (i == '-' && s == "") {
			s += string(s)
		}
	}
	a, _ := strconv.Atoi(s)
}
