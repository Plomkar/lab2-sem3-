package tree

import "fmt"

// CBTree — полное бинарное дерево в массиве (дети i — 2i+1, 2i+2).
type CBTree struct {
	Data []int
}

func (t *CBTree) Init() { t.Data = t.Data[:0] }

func (t *CBTree) Insert(value int) { t.Data = append(t.Data, value) }

func (t *CBTree) Get(index int) int {
	if index < 0 || index >= len(t.Data) {
		return -1
	}
	return t.Data[index]
}

func (t *CBTree) Search(value int) int {
	for i, v := range t.Data {
		if v == value {
			return i
		}
	}
	return -1
}

func (t *CBTree) CheckComplete() bool {
	n := len(t.Data)
	if n == 0 {
		return true
	}
	sawNull := false
	for i := 0; i < n; i++ {
		left, right := 2*i+1, 2*i+2
		if left < n && sawNull {
			return false
		}
		if left >= n {
			sawNull = true
		}
		if right < n && sawNull {
			return false
		}
		if right >= n {
			sawNull = true
		}
	}
	return true
}

func (t *CBTree) Print() {
	for _, v := range t.Data {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
