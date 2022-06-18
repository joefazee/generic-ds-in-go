package list

type node[T any] struct {
	value      T
	next, prev *node[T]
}

func (n node[T]) Value() T {
	return n.value
}

func (n node[T]) Next() *node[T] {
	return n.next
}

func (n node[T]) Prev() *node[T] {
	return n.prev
}

type linkedList[T any] struct {
	head, tail *node[T]
}

func New[T any]() *linkedList[T] {
	return &linkedList[T]{}
}

func (l *linkedList[T]) PushFront(value T) {

	newNode := &node[T]{value: value}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		oldNode := l.head
		oldNode.prev = newNode
		newNode.next = oldNode
		l.head = newNode
	}
}

func (l *linkedList[T]) PushBack(value T) {

	newNode := &node[T]{value: value}
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		oldNode := l.tail
		newNode.prev = oldNode
		oldNode.next = newNode
		l.tail = newNode
	}
}

func (l *linkedList[T]) Head() *node[T] {
	return l.head
}

func (l *linkedList[T]) Tail() *node[T] {
	return l.tail
}
