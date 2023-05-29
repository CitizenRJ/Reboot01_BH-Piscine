package piscine

func Split(s, sep string) []string {
	str := ""
	line := []string{}
	for i := 0, i < len (s)-len((sep)-1); i++ {
		if string(s[i:i+len(sep)]) == sep {
			line = append(line, str)
			str = ""
			i = i + (len(sep)-1)
		} else {
			str = str + string(s[i])
		}
	}
	str = str + string(s[len(s)-(len(sep)-1):])
	line = append(line, str)
	return line
}