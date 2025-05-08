package main

import "fmt"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

func Empty[T comparable]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Add(t T) {
	node := &Node[T]{
		Value: t,
	}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	l.Tail.Next = node
	l.Tail = l.Tail.Next
}

func (l *List[T]) Insert(t T, position int) {
	node := &Node[T]{
		Value: t,
	}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}

	if position <= 0 {
		node.Next = l.Head
		l.Head = node
		return
	}

	current := l.Head

	for i := 1; i < position; i++ {
		if current.Next == nil {
			current.Next = node
			l.Tail = current.Next
			return
		}

		current = current.Next
	}

	node.Next = current.Next
	current.Next = node

	if l.Tail == current {
		l.Tail = node
	}
}

func (l *List[T]) Index(t T) int {
	i := 0

	for current := l.Head; current != nil; current = current.Next {
		if current.Value == t {
			return i
		}

		i += 1
	}

	return -1
}

func main() {
	list := Empty[int]()

	list.Add(5)
	list.Add(10)

	fmt.Println(list.Index(5))
	fmt.Println(list.Index(10))
	fmt.Println(list.Index(20))

	list.Insert(100, 0)

	fmt.Println(list.Index(5))
	fmt.Println(list.Index(10))
	fmt.Println(list.Index(20))
	fmt.Println(list.Index(100))

	list.Insert(200, 1)

	fmt.Println(list.Index(5))
	fmt.Println(list.Index(10))
	fmt.Println(list.Index(20))
	fmt.Println(list.Index(100))
	fmt.Println(list.Index(200))
}
