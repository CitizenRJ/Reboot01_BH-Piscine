package piscine

func ConcatParams(args []string) string {
	str := ""
	for _, arg := range args {
		str += arg + "\n"
	}
	return str
}
