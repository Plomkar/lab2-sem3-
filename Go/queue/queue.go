package queue

import "fmt"

type QNode struct {
	Data int
	Next *QNode
}

type Queue struct {
	Head *QNode
	Tail *QNode
}

func (q *Queue) Init() { q.Head, q.Tail = nil, nil }

func (q *Queue) IsEmpty() bool { return q.Head == nil }

func (q *Queue) Push(value int) {
	n := &QNode{Data: value}
	if q.Tail != nil {
		q.Tail.Next = n
	} else {
		q.Head = n
	}
	q.Tail = n
}

func (q *Queue) Pop() int {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return -1
	}
	del := q.Head
	val := del.Data
	q.Head = del.Next
	if q.Head == nil {
		q.Tail = nil
	}
	return val
}

func (q *Queue) Peek() int {
	if q.IsEmpty() {
		return -1
	}
	return q.Head.Data
}

func (q *Queue) Print() {
	cur := q.Head
	for cur != nil {
		fmt.Printf("%d ", cur.Data)
		cur = cur.Next
	}
	fmt.Println()
}
