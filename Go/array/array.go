package array

type Array struct {
	Data []int
}

func New() *Array { return &Array{} }

func (a *Array) Init() { a.Data = a.Data[:0] }

func (a *Array) Push(value int) { a.Data = append(a.Data, value) }

func (a *Array) PushAt(index, value int) {
	if index < 0 || index > len(a.Data) {
		return
	}
	a.Data = append(a.Data, 0)
	copy(a.Data[index+1:], a.Data[index:])
	a.Data[index] = value
}

func (a *Array) Get(index int) int {
	if index < 0 || index >= len(a.Data) {
		return -1
	}
	return a.Data[index]
}

func (a *Array) Del(index int) {
	if index < 0 || index >= len(a.Data) {
		return
	}
	a.Data = append(a.Data[:index], a.Data[index+1:]...)
}

func (a *Array) Set(index, value int) {
	if index < 0 || index >= len(a.Data) {
		return
	}
	a.Data[index] = value
}

func (a *Array) Len() int { return len(a.Data) }
