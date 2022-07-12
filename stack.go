package main

import (
	"errors"
	"fmt"
	"log"
)

type stack struct {
	top *node
}

type node struct {
	val  int
	next *node
}

func newStack() *stack {
	return &stack{
		top: nil,
	}
}

//LIFO Last In First Out

func (s *stack) push(val int) {

	if s.top == nil {
		n := &node{val: val, next: nil}
		s.top = n
		return
	}

	n := &node{val: val, next: s.top}
	s.top = n
}

func (s stack) peek() int {
	return s.top.val
}

func (s *stack) pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}

	t := s.top
	s.top = s.top.next
	return t.val, nil
}

func (s stack) isEmpty() bool {
	return s.top == nil
}

func main() {
	s := newStack()

	s.push(1)
	s.push(3)
	s.push(4)
	s.push(5)

	for !s.isEmpty() {
		v, err := s.pop()
		checkStackErr(err)
		fmt.Println("pop :", v)
	}

}

func checkStackErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
