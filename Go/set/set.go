package set

import (
	"project/array"
	"project/flist"
)

// Set — АТД «Множество» на хеш-таблице с цепочками.
type Set struct {
	Buckets []flist.FList
	Size    int
}

func (s *Set) Init(capacity int) {
	if capacity < 4 {
		capacity = 4
	}
	s.Buckets = make([]flist.FList, capacity)
	for i := range s.Buckets {
		s.Buckets[i].Init()
	}
	s.Size = 0
}

func (s *Set) Clear() {
	for i := range s.Buckets {
		for s.Buckets[i].Head != nil {
			s.Buckets[i].Del(0)
		}
	}
	s.Size = 0
}

func (s *Set) bucketOf(value int) int {
	cap := len(s.Buckets)
	r := value % cap
	if r < 0 {
		r += cap
	}
	return r
}

func (s *Set) rehash() {
	old := s.Buckets
	newCap := len(old) * 2
	s.Buckets = make([]flist.FList, newCap)
	for i := range s.Buckets {
		s.Buckets[i].Init()
	}
	for i := range old {
		cur := old[i].Head
		for cur != nil {
			next := cur.Next
			idx := s.bucketOf(cur.Data)
			cur.Next = s.Buckets[idx].Head
			s.Buckets[idx].Head = cur
			cur = next
		}
	}
}

func (s *Set) At(value int) bool {
	return s.Buckets[s.bucketOf(value)].Search(value) != -1
}

func (s *Set) Add(value int) bool {
	if s.At(value) {
		return false
	}
	if s.Size+1 > len(s.Buckets) {
		s.rehash()
	}
	s.Buckets[s.bucketOf(value)].Push(value)
	s.Size++
	return true
}

func (s *Set) Del(value int) bool {
	b := &s.Buckets[s.bucketOf(value)]
	if b.Search(value) == -1 {
		return false
	}
	b.DelValue(value)
	s.Size--
	return true
}

func (s *Set) Len() int { return s.Size }

func (s *Set) ToArray(out *array.Array) {
	out.Init()
	for i := range s.Buckets {
		for cur := s.Buckets[i].Head; cur != nil; cur = cur.Next {
			out.Push(cur.Data)
		}
	}
}
