package queue

import (
	"bytes"
	"errors"
	"fmt"
)

type q[T any] struct {
	items []T
	back  int
	front int
}

func New[T any]() *q[T] {
	return &q[T]{
		items: make([]T, 10),
		front: 0,
		back:  -1,
	}
}

func (this *q[T]) Enqueue(item T) *q[T] {
	if this.back == len(this.items)-1 {
		newList := make([]T, len(this.items)*2)
		copy(newList, this.items)
		this.items = newList
	}

	this.back++
	this.items[this.back] = item

	return this
}

func (this *q[T]) Dequeue() (T, error) {
	if this.back == -1 || this.front > this.back {
		var res T
		return res, errors.New("empty queue")
	}

	res := this.items[this.front]
	this.front++
	return res, nil
}

func (this *q[T]) String() string {

	var buf bytes.Buffer
	for {
		res, err := this.Dequeue()
		if err != nil {
			break
		}

		buf.WriteString(fmt.Sprintf(" %v ", res))
	}

	return buf.String()
}
