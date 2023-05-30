package piscine

func CountIf(f func(string) bool, tab []string) int {

	result := 0

	for _, value := range tab {
		if f(value) {
			result ++
		}
	}
	return result
}

