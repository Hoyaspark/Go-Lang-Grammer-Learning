package main

import "fmt"

type queue struct {
	item chan interface{}
}

func (q *queue) push(val any) {
	q.item <- val
}

func (q *queue) pop() any {
	return <-q.item
}

func NewQueue(size int) *queue {
	return &queue{item: make(chan interface{}, size)}
}

func main() {
	q := NewQueue(3)

	q.push(3)
	q.push("hello")
	q.push(4)
	q.push(5)

	// stop? yes..
	fmt.Println(q.pop())
	fmt.Println(q.pop())
	fmt.Println(q.pop())

}
