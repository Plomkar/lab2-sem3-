package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"project/array"
	"project/set"
)

var (
	gVals  array.Array
	gGroup array.Array
	gUsed  set.Set
	gN     int
	gS     int
	gK     int
)

func dfs(start, cur, idx int) bool {
	if idx == gK {
		return true
	}
	if cur == gS {
		return dfs(0, 0, idx+1)
	}
	for i := start; i < gN; i++ {
		if gUsed.At(i) {
			continue
		}
		v := gVals.Get(i)
		if cur+v > gS {
			continue
		}
		gUsed.Add(i)
		gGroup.Set(i, idx)
		if dfs(i+1, cur+v, idx) {
			return true
		}
		gUsed.Del(i)
		if cur == 0 {
			break
		}
	}
	return false
}

func tryPartition(total int64, S int, print bool) bool {
	maxV := gVals.Get(0)
	if S < maxV || total%int64(S) != 0 {
		return false
	}
	gK = int(total / int64(S))
	gS = S
	if gK < 2 {
		return false
	}
	gUsed.Init(16)
	gGroup.Init()
	for i := 0; i < gN; i++ {
		gGroup.Push(-1)
	}
	ok := dfs(0, 0, 0)
	if ok && print {
		fmt.Printf("Сумма каждого подмножества = %d, подмножеств: %d\n", S, gK)
		for j := 0; j < gK; j++ {
			fmt.Print("{")
			first := true
			for i := 0; i < gN; i++ {
				if gGroup.Get(i) == j {
					if !first {
						fmt.Print(", ")
					}
					fmt.Printf("%d", gVals.Get(i))
					first = false
				}
			}
			fmt.Println("}")
		}
	}
	gUsed.Clear()
	return ok
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	if !sc.Scan() {
		fmt.Println("Ошибка: некорректный ввод")
		return
	}
	line := strings.TrimSpace(sc.Text())
	fields := strings.Fields(line)
	if len(fields) == 0 {
		fmt.Println("Ошибка: некорректный ввод")
		return
	}
	n, err := strconv.Atoi(fields[0])
	if err != nil || n <= 0 {
		fmt.Println("Ошибка: некорректный ввод")
		return
	}
	if len(fields) < 1+n {
		fmt.Println("Ошибка: некорректный ввод")
		return
	}

	var s set.Set
	s.Init(16)
	for i := 0; i < n; i++ {
		x, err := strconv.ParseInt(fields[1+i], 10, 64)
		if err != nil || x <= 0 {
			fmt.Println("Ошибка: нужны натуральные числа")
			return
		}
		s.Add(int(x))
	}

	S := -1
	if len(fields) > 1+n {
		if tmp, err := strconv.Atoi(fields[1+n]); err == nil {
			if tmp <= 0 {
				fmt.Println("Ошибка: сумма S должна быть натуральным числом")
				return
			}
			S = tmp
		}
	}

	s.ToArray(&gVals)
	sort.Slice(gVals.Data, func(i, j int) bool {
		return gVals.Data[i] > gVals.Data[j]
	})
	gN = gVals.Len()
	var total int64
	for i := 0; i < gN; i++ {
		total += int64(gVals.Get(i))
	}

	any := false
	if S > 0 {
		any = tryPartition(total, S, true)
	} else {
		maxV := gVals.Get(0)
		for cand := int64(maxV); cand <= total/2; cand++ {
			if total%cand == 0 {
				if tryPartition(total, int(cand), true) {
					any = true
				}
			}
		}
	}
	if !any {
		fmt.Println("Ошибка: множество нельзя разбить на подмножества с равной суммой")
	}
	s.Clear()
}
