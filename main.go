package main

import (
	"fmt"
	"generics-examples/list"
	"generics-examples/queue"
)

type user struct {
	name string
}

func main() {

	q := queue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(10)

	fmt.Println(q)

	// ==================== LINKEDLIST EXAMPLE =================

	l := list.New[string]()
	l.PushFront("A")
	l.PushFront("B")
	l.PushFront("C")
	l.PushFront("D")

	node := l.Head()
	for node != nil {
		fmt.Println(node.Value())
		node = node.Next()
	}

}
