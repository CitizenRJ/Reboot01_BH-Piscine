package piscine

func Compact(ptr *[]string) int {
	count := 0
	var temp []string
	for i := 0; i < len(*ptr); i++ {
		if (*prt)[i] != "" {
			temp = append(temp, (*prt)[i])
			count++
		}
	}
	*ptr = temp
	return count
}
