package list

import "fmt"

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type LList struct {
	Head *Node
	Tail *Node
}

func (l *LList) Init() { l.Head, l.Tail = nil, nil }

func (l *LList) Push(value int) {
	n := &Node{Data: value, Next: l.Head}
	if l.Head != nil {
		l.Head.Prev = n
	} else {
		l.Tail = n
	}
	l.Head = n
}

func (l *LList) PushEnd(value int) {
	n := &Node{Data: value, Prev: l.Tail}
	if l.Tail != nil {
		l.Tail.Next = n
	} else {
		l.Head = n
	}
	l.Tail = n
}

func (l *LList) PushAt(index, value int) {
	if index <= 0 {
		l.Push(value)
		return
	}
	cur := l.Head
	for i := 0; i < index-1 && cur != nil; i++ {
		cur = cur.Next
	}
	if cur == nil {
		l.PushEnd(value)
		return
	}
	n := &Node{Data: value, Prev: cur, Next: cur.Next}
	if cur.Next != nil {
		cur.Next.Prev = n
	} else {
		l.Tail = n
	}
	cur.Next = n
}

func (l *LList) PushAfter(index, value int) {
	cur := l.Head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return
	}
	n := &Node{Data: value, Prev: cur, Next: cur.Next}
	if cur.Next != nil {
		cur.Next.Prev = n
	} else {
		l.Tail = n
	}
	cur.Next = n
}

func (l *LList) Del(index int) {
	cur := l.Head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return
	}
	l.unlink(cur)
}

func (l *LList) DelValue(value int) {
	cur := l.Head
	for cur != nil && cur.Data != value {
		cur = cur.Next
	}
	if cur == nil {
		return
	}
	l.unlink(cur)
}

func (l *LList) unlink(cur *Node) {
	if cur.Prev != nil {
		cur.Prev.Next = cur.Next
	} else {
		l.Head = cur.Next
	}
	if cur.Next != nil {
		cur.Next.Prev = cur.Prev
	} else {
		l.Tail = cur.Prev
	}
}

func (l *LList) Get(index int) int {
	cur := l.Head
	for i := 0; i < index && cur != nil; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return -1
	}
	return cur.Data
}

func (l *LList) Search(value int) int {
	cur := l.Head
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

func (l *LList) Print() {
	cur := l.Head
	for cur != nil {
		fmt.Printf("%d ", cur.Data)
		cur = cur.Next
	}
	fmt.Println()
}

func (l *LList) PrintReverse() {
	cur := l.Tail
	for cur != nil {
		fmt.Printf("%d ", cur.Data)
		cur = cur.Prev
	}
	fmt.Println()
}
