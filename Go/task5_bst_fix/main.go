package main

import (
	"bufio"
	"fmt"
	"os"

	"project/array"
	"project/stack"
	"project/tree"
)

func inorder(t *tree.CBTree, order *array.Array) {
	order.Init()
	n := len(t.Data)
	var st stack.Stack
	st.Init()
	cur := 0
	for cur < n || !st.IsEmpty() {
		for cur < n {
			st.Push(cur)
			cur = 2*cur + 1
		}
		cur = st.Pop()
		order.Push(cur)
		cur = 2*cur + 2
	}
}

func isSorted(t *tree.CBTree, order *array.Array) bool {
	for i := 1; i < order.Len(); i++ {
		if t.Get(order.Get(i-1)) >= t.Get(order.Get(i)) {
			return false
		}
	}
	return true
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var n int
	if !sc.Scan() {
		fmt.Println("Ошибка: некорректный ввод")
		return
	}
	fmt.Sscan(sc.Text(), &n)
	if n <= 0 {
		fmt.Println("Ошибка: некорректный ввод")
		return
	}
	var t tree.CBTree
	t.Init()
	for i := 0; i < n; i++ {
		if !sc.Scan() {
			fmt.Printf("Ошибка: ожидалось %d чисел\n", n)
			return
		}
		var v int
		fmt.Sscan(sc.Text(), &v)
		t.Insert(v)
	}

	var order array.Array
	inorder(&t, &order)

	first, second := -1, -1
	for i := 1; i < order.Len(); i++ {
		a, b := order.Get(i-1), order.Get(i)
		if t.Get(a) > t.Get(b) {
			if first == -1 {
				first = a
			}
			second = b
		}
	}

	if first == -1 {
		fmt.Println("Дерево уже является корректным BST:")
		t.Print()
		return
	}

	t.Data[first], t.Data[second] = t.Data[second], t.Data[first]

	if !isSorted(&t, &order) {
		fmt.Println("Ошибка: дерево нельзя исправить обменом двух узлов")
		return
	}
	fmt.Printf("Поменяны местами узлы с индексами %d и %d\n", first, second)
	fmt.Print("Исправленное дерево (по уровням): ")
	t.Print()
}
