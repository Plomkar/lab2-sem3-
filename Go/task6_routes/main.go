package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"project/array"
	"project/flist"
	"project/hashtable"
	"project/queue"
	"project/stack"
)

const INF = 1000000000

type Network struct {
	IDs   *hashtable.HashTable
	Delay *hashtable.HashTable
	Names []string
	Adj   []flist.FList
}

var net Network

func getId(name string, create bool) int {
	if v, ok := net.IDs.Get(name); ok {
		return v.(int)
	}
	if !create {
		return -1
	}
	id := len(net.Names)
	net.Names = append(net.Names, name)
	var f flist.FList
	f.Init()
	net.Adj = append(net.Adj, f)
	net.IDs.Put(name, id)
	return id
}

func addRoute(a, b string, d int) {
	u, v := getId(a, true), getId(b, true)
	exists := net.Delay.Has(a + "|" + b)
	net.Delay.Put(a+"|"+b, d)
	net.Delay.Put(b+"|"+a, d)
	if !exists && u != v {
		net.Adj[u].Push(v)
		net.Adj[v].Push(u)
	}
}

func findPath(from, to string) {
	s, t := getId(from, false), getId(to, false)
	if s == -1 || t == -1 {
		fmt.Println("Сервер не найден")
		return
	}

	n := len(net.Names)
	var dist, parent, inQ array.Array
	dist.Init()
	parent.Init()
	inQ.Init()
	for i := 0; i < n; i++ {
		dist.Push(INF)
		parent.Push(-1)
		inQ.Push(0)
	}

	var q queue.Queue
	q.Init()
	dist.Set(s, 0)
	q.Push(s)
	inQ.Set(s, 1)
	for !q.IsEmpty() {
		u := q.Pop()
		inQ.Set(u, 0)
		for e := net.Adj[u].Head; e != nil; e = e.Next {
			v := e.Data
			wAny, _ := net.Delay.Get(net.Names[u] + "|" + net.Names[v])
			w := wAny.(int)
			if dist.Get(u)+w < dist.Get(v) {
				dist.Set(v, dist.Get(u)+w)
				parent.Set(v, u)
				if inQ.Get(v) == 0 {
					q.Push(v)
					inQ.Set(v, 1)
				}
			}
		}
	}

	if dist.Get(t) == INF {
		fmt.Println("Маршрута между серверами нет")
		return
	}

	var path stack.Stack
	path.Init()
	for v := t; v != -1; v = parent.Get(v) {
		path.Push(v)
	}
	fmt.Printf("минимальная задержка = %d (маршрут ", dist.Get(t))
	first := true
	for !path.IsEmpty() {
		if !first {
			fmt.Print(" -> ")
		}
		fmt.Printf("\"%s\"", net.Names[path.Pop()])
		first = false
	}
	fmt.Println(")")
}

func main() {
	net.IDs = hashtable.New(hashtable.DoubleHashing, 11, true)
	net.Delay = hashtable.New(hashtable.DoubleHashing, 11, true)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		if len(fields) == 0 {
			continue
		}
		cmd := fields[0]
		switch cmd {
		case "ADD", "UPDATE":
			if len(fields) >= 4 {
				d, err := strconv.Atoi(fields[3])
				if err == nil && d >= 0 {
					addRoute(fields[1], fields[2], d)
					fmt.Println("OK")
					continue
				}
			}
			fmt.Printf("Ошибка: ожидалось  %s A B d  (d >= 0)\n", cmd)
		case "PATH":
			if len(fields) >= 3 {
				findPath(fields[1], fields[2])
			} else {
				fmt.Println("Ошибка: ожидалось  PATH A B")
			}
		case "EXIT":
			return
		default:
			fmt.Println("Unknown command:", cmd)
		}
	}
}
