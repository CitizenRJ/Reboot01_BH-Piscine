package piscine

func Split(s, sep string) []string {
	line := []string{}
	var str string
	for _, c := range s {
		if string(c) != sep {
			str = str + string(c)
		} else {
			line = append(line, str)
			str = ""
		}
	}
	line = append(line, str)
	return line
}
