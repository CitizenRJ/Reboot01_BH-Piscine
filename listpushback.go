package piscine

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		return
	}

	listItem := l.Head
	for listItem.Next != nil {
		listItem = listItem.Next
	}
	listItem.Next = n
}
