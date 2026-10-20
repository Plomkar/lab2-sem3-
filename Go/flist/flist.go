package flist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type FList struct {
	Head *Node
}

func (f *FList) Init() { f.Head = nil }

func (f *FList) Push(value int) {
	f.Head = &Node{Data: value, Next: f.Head}
}

func (f *FList) PushEnd(value int) {
	newNode := &Node{Data: value}
	if f.Head == nil {
		f.Head = newNode
		return
	}
	cur := f.Head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = newNode
}

func (f *FList) PushAt(index, value int) {
	if index <= 0 {
		f.Push(value)
		return
	}
	cur := f.Head
	for i := 0; i < index-1 && cur != nil; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return
	}
	cur.Next = &Node{Data: value, Next: cur.Next}
}

func (f *FList) PushAfter(index, value int) {
	cur := f.Head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return
	}
	cur.Next = &Node{Data: value, Next: cur.Next}
}

func (f *FList) Del(index int) {
	if f.Head == nil {
		return
	}
	if index == 0 {
		f.Head = f.Head.Next
		return
	}
	cur := f.Head
	for i := 0; i < index-1 && cur.Next != nil; i++ {
		cur = cur.Next
	}
	if cur.Next == nil {
		return
	}
	cur.Next = cur.Next.Next
}

func (f *FList) DelValue(value int) {
	if f.Head == nil {
		return
	}
	if f.Head.Data == value {
		f.Head = f.Head.Next
		return
	}
	cur := f.Head
	for cur.Next != nil && cur.Next.Data != value {
		cur = cur.Next
	}
	if cur.Next == nil {
		return
	}
	cur.Next = cur.Next.Next
}

func (f *FList) Get(index int) int {
	cur := f.Head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return -1
	}
	return cur.Data
}

func (f *FList) Search(value int) int {
	cur := f.Head
	idx := 0
	for cur != nil {
		if cur.Data == value {
			return idx
		}
		cur = cur.Next
		idx++
	}
	return -1
}

func (f *FList) Print() {
	cur := f.Head
	for cur != nil {
		fmt.Printf("%d ", cur.Data)
		cur = cur.Next
	}
	fmt.Println()
}
