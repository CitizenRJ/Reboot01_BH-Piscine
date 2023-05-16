package piscine

func StrRev(s string) string {
	var a string
	byt := []byte(s)
	for _, word := range byt {
		a = string(word) + a
	}
	return a
}
