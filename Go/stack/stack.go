package stack

import "fmt"

type SNode struct {
	Data int
	Next *SNode
}

type Stack struct {
	Head *SNode
}

func (s *Stack) Init() { s.Head = nil }

func (s *Stack) IsEmpty() bool { return s.Head == nil }

func (s *Stack) Push(value int) {
	s.Head = &SNode{Data: value, Next: s.Head}
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return -1
	}
	del := s.Head
	val := del.Data
	s.Head = del.Next
	return val
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.Head.Data
}

func (s *Stack) Print() {
	cur := s.Head
	for cur != nil {
		fmt.Printf("%d ", cur.Data)
		cur = cur.Next
	}
	fmt.Println()
}
