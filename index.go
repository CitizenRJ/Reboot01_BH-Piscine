package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == (toFind)[0] {
			for j := 0; j < len(toFind); j++ {
				if s[i+j] != toFind[j] {
					break
				}
			}
			if j == len(toFind) {
				return 1
			}
		}
	}

	return -1
}
