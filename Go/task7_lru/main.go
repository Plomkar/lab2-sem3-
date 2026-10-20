package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"project/hashtable"
	"project/list"
)

type CacheItem struct {
	Value int
	Node  *list.Node
}

func unlink(l *list.LList, n *list.Node) {
	if n.Prev != nil {
		n.Prev.Next = n.Next
	} else {
		l.Head = n.Next
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
	} else {
		l.Tail = n.Prev
	}
	n.Prev, n.Next = nil, nil
}

func pushFront(l *list.LList, n *list.Node) {
	n.Prev = nil
	n.Next = l.Head
	if l.Head != nil {
		l.Head.Prev = n
	} else {
		l.Tail = n
	}
	l.Head = n
}

type LRUCache struct {
	cap, size int
	order     list.LList
	table     *hashtable.HashTable
}

func NewLRUCache(capacity int) *LRUCache {
	c := &LRUCache{cap: capacity, table: hashtable.New(hashtable.DoubleHashing, 11, true)}
	c.order.Init()
	return c
}

func (c *LRUCache) Get(key int) int {
	v, ok := c.table.Get(strconv.Itoa(key))
	if !ok {
		return -1
	}
	it := v.(*CacheItem)
	unlink(&c.order, it.Node)
	pushFront(&c.order, it.Node)
	return it.Value
}

func (c *LRUCache) Set(key, value int) {
	k := strconv.Itoa(key)
	if v, ok := c.table.Get(k); ok {
		it := v.(*CacheItem)
		it.Value = value
		unlink(&c.order, it.Node)
		pushFront(&c.order, it.Node)
		return
	}
	if c.cap <= 0 {
		return
	}
	if c.size == c.cap {
		old := c.order.Tail
		c.table.Del(strconv.Itoa(old.Data))
		unlink(&c.order, old)
		c.size--
	}
	c.order.Push(key)
	item := &CacheItem{Value: value, Node: c.order.Head}
	c.table.Put(k, item)
	c.size++
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	if !sc.Scan() {
		fmt.Println("Ошибка ввода: ожидалось  cap Q")
		return
	}
	cap, _ := strconv.Atoi(sc.Text())
	if !sc.Scan() {
		fmt.Println("Ошибка ввода: ожидалось  cap Q")
		return
	}
	q, _ := strconv.Atoi(sc.Text())

	cache := NewLRUCache(cap)
	first := true
	for i := 0; i < q && sc.Scan(); i++ {
		cmd := sc.Text()
		switch cmd {
		case "SET":
			if !sc.Scan() {
				return
			}
			x, _ := strconv.Atoi(sc.Text())
			if !sc.Scan() {
				return
			}
			y, _ := strconv.Atoi(sc.Text())
			cache.Set(x, y)
		case "GET":
			if !sc.Scan() {
				return
			}
			x, _ := strconv.Atoi(sc.Text())
			if !first {
				fmt.Print(" ")
			}
			fmt.Print(cache.Get(x))
			first = false
		default:
			fmt.Println("\nНеизвестный запрос:", cmd)
			return
		}
	}
	fmt.Println()
}
