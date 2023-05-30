package piscine

func Split(s, sep string) []string {
	a := []string{}
	b := ""
	length := 0
	for i, v := range s {
		if len(s) == i+1 {
			b += string(v)
			a = append(a, b)
		}
		if len(sep) == length {
			a = append(a, b[0:len(b)-len(sep)])
			b = ""
			length = 0
		}
		if string(v) == string(sep[length]) {
			length++
			b += string(v)
		} else {
			b += string(v)
			length = 0
			if string(v) == string(sep[length]) {
				length++
			}
		}
	}
	return a
}
