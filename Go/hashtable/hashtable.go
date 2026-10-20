package hashtable

import "math/big"

// HashMode — режим разрешения коллизий.
type HashMode int

const (
	LinearProbing HashMode = iota
	DoubleHashing
)

// SlotState — состояние слота.
type SlotState int

const (
	Empty SlotState = iota
	Used
	Deleted
)

// Entry — запись таблицы. Value хранится как any.
type Entry struct {
	Key   string
	Value any
	State SlotState
}

// HashTable — хеш-таблица с открытой адресацией.
type HashTable struct {
	Slots      []Entry
	Live       int
	Filled     int
	Mode       HashMode
	AutoResize bool
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func NextPrime(n int) int {
	for !IsPrime(n) {
		n++
	}
	return n
}

// Hash1 — полиномиальная хеш-функция.
func Hash1(s string) uint64 {
	var h uint64
	for i := 0; i < len(s); i++ {
		h = h*131 + uint64(s[i])
	}
	return h
}

// Hash2 — FNV-1a.
func Hash2(s string) uint64 {
	h := uint64(1469598103934665603)
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= 1099511628211
	}
	return h
}

func New(mode HashMode, capacity int, autoResize bool) *HashTable {
	t := &HashTable{
		Slots:      make([]Entry, NextPrime(capacity)),
		Mode:       mode,
		AutoResize: autoResize,
	}
	for i := range t.Slots {
		t.Slots[i].State = Empty
	}
	return t
}

func (t *HashTable) Size() int { return t.Live }

// probe — i-я проба для ключа.
func (t *HashTable) probe(key string, i int) int {
	cap := len(t.Slots)
	h1 := int(Hash1(key) % uint64(cap))
	step := 1
	if t.Mode == DoubleHashing {
		step = 1 + int(Hash2(key)%uint64(cap-1))
	}
	return (h1 + i*step) % cap
}

// Get возвращает значение или nil.
func (t *HashTable) Get(key string) (any, bool) {
	cap := len(t.Slots)
	for i := 0; i < cap; i++ {
		e := &t.Slots[t.probe(key, i)]
		if e.State == Empty {
			return nil, false
		}
		if e.State == Used && e.Key == key {
			return e.Value, true
		}
	}
	return nil, false
}

func (t *HashTable) Has(key string) bool {
	_, ok := t.Get(key)
	return ok
}

func (t *HashTable) rehash(newCap int) {
	old := t.Slots
	t.Slots = make([]Entry, NextPrime(newCap))
	for i := range t.Slots {
		t.Slots[i].State = Empty
	}
	t.Live, t.Filled = 0, 0
	savedMode, savedAuto := t.Mode, t.AutoResize
	t.AutoResize = false
	for _, e := range old {
		if e.State == Used {
			cap := len(t.Slots)
			for i := 0; i < cap; i++ {
				d := &t.Slots[t.probe(e.Key, i)]
				if d.State != Used {
					d.Key, d.Value, d.State = e.Key, e.Value, Used
					t.Live++
					t.Filled++
					break
				}
			}
		}
	}
	t.Mode, t.AutoResize = savedMode, savedAuto
}

// Put — вставка или обновление.
func (t *HashTable) Put(key string, value any) bool {
	if t.AutoResize && (t.Filled+1)*10 > len(t.Slots)*7 {
		t.rehash(len(t.Slots) * 2)
	}
	cap := len(t.Slots)
	firstFree := -1
	for i := 0; i < cap; i++ {
		idx := t.probe(key, i)
		e := &t.Slots[idx]
		if e.State == Used && e.Key == key {
			e.Value = value
			return true
		}
		if e.State == Deleted && firstFree == -1 {
			firstFree = idx
		}
		if e.State == Empty {
			if firstFree == -1 {
				firstFree = idx
			}
			break
		}
	}
	if firstFree == -1 {
		return false
	}
	e := &t.Slots[firstFree]
	if e.State == Empty {
		t.Filled++
	}
	e.Key, e.Value, e.State = key, value, Used
	t.Live++
	return true
}

func (t *HashTable) Del(key string) bool {
	cap := len(t.Slots)
	for i := 0; i < cap; i++ {
		e := &t.Slots[t.probe(key, i)]
		if e.State == Empty {
			return false
		}
		if e.State == Used && e.Key == key {
			e.State = Deleted
			t.Live--
			return true
		}
	}
	return false
}

// вспомогательное для совместимости с big (не используется, но оставлено для ясности)
var _ = big.NewInt
