package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	size int
	head *ListItem
	tail *ListItem
}

func (l list) Len() int {
	return l.size
}

func (l list) Front() *ListItem {
	return l.head
}

func (l list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	i := &ListItem{v, nil, nil}
	if l.Len() >= 1 {
		i.Prev, l.head.Next = l.head, i
	} else {
		l.tail = i
	}
	l.head = i
	l.size++

	return i
}

func (l *list) PushBack(v interface{}) *ListItem {
	i := &ListItem{v, nil, nil}
	if l.Len() >= 1 {
		i.Next, l.tail.Prev = l.tail, i
	} else {
		l.head = i
	}
	l.tail = i
	l.size++

	return i
}

func (l *list) Remove(i *ListItem) {
	switch {
	case i.Next == nil && i.Prev == nil:
		l.head, l.tail = nil, nil
	case i.Next == nil:
		l.head = i.Prev
		i.Prev.Next = nil
	case i.Prev == nil:
		l.tail = i.Next
		i.Next.Prev = nil
	default:
		i.Next.Prev = i.Prev
		i.Prev.Next = i.Next
	}
	l.size--
}

func (l *list) MoveToFront(i *ListItem) {

}

func NewList() List {
	return new(list)
}
