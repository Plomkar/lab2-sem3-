package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"project/array"
	"project/hashtable"
	"project/set"
)

var (
	gIDs   array.Array
	gNames []string
	gOut   []string
)

func emit(cur *array.Array) {
	var sb strings.Builder
	sb.WriteString("{")
	for i := 0; i < cur.Len(); i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(gNames[cur.Get(i)])
	}
	sb.WriteString("}")
	gOut = append(gOut, sb.String())
}

func gen(start, need int, cur *array.Array) {
	if need == 0 {
		emit(cur)
		return
	}
	n := gIDs.Len()
	var tried set.Set
	tried.Init(8)
	for i := start; i+need <= n; i++ {
		id := gIDs.Get(i)
		if tried.At(id) {
			continue
		}
		tried.Add(id)
		cur.Push(id)
		gen(i+1, need-1, cur)
		cur.Del(cur.Len() - 1)
	}
	tried.Clear()
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fields := strings.Fields(sc.Text())

	dict := hashtable.New(hashtable.DoubleHashing, 11, true)
	gIDs.Init()
	for _, tok := range fields {
		if v, ok := dict.Get(tok); ok {
			gIDs.Push(v.(int))
		} else {
			id := len(gNames)
			gNames = append(gNames, tok)
			dict.Put(tok, id)
			gIDs.Push(id)
		}
	}
	n := gIDs.Len()
	if n == 0 {
		fmt.Println("[{}]")
		return
	}
	if n > 20 {
		fmt.Println("Слишком длинный массив (результат содержит до 2^n подмассивов), максимум 20 элементов")
		return
	}

	var cur array.Array
	cur.Init()
	for l := 0; l <= n; l++ {
		gen(0, l, &cur)
	}
	fmt.Print("[")
	for i, s := range gOut {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(s)
	}
	fmt.Println("]")
}
