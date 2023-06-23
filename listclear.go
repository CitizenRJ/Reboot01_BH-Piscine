package piscine

func ListClear(l *List) {
	tmp := l.Head
	next := l.Head
	for tmp != nil {
		next = tmp.Next
		tmp = next
	}
	l.Head = nil
}
