package piscine

func Split(s, sep string) []string {
	results := make([]string, 0)
	for _, i := range s {
		j := string(i)
		if j != sep {
			results = append(results, string(i))
		}
	}
	return results
}
