package main

import (
	"fmt"
	"bufio"
	"container/list"
	"os"
)

type Queue[T any] struct {
	items *list.List
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: list.New()}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items.PushBack(item)
}

func (q  Queue[T]) Dequeue() T {
	if q.items.Len() == 0 {
		var zero T
		return zero
	}
	front := q.items.Front().Value.(T)
	q.items.Remove(q.items.Front())
	return front
}

func (q *Queue[T]) Size() int {
	return q.items.Len()
}


func main() {

	queue := NewQueue[string]()

	for c:= 'A'; c <= 'P'; c++ {
		queue.Enqueue(string(c))
	}

	scanner := bufio.NewScanner(os.Stdin)
	jogosContador := 1

	for queue.Size() > 1 {
		team.Left := queue.Dequeue()
		team.Right := queue.Dequeue()
		
	}
}
