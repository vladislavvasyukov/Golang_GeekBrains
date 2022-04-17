package main

import (
	"sync"
	"testing"
)

type Set struct {
	sync.Mutex
	mm map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		mm: map[int]struct{}{},
	}
}

func (s *Set) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *Set) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

func SetAdd(b *testing.B, p int) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(p)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
	})
}

func SetHas(b *testing.B, p int) {
	var set = NewSet()
	b.Run("", func(b *testing.B) {
		b.SetParallelism(p)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkSetAdd100(b *testing.B) {
	SetAdd(b, 100)
}

func BenchmarkSetHas900(b *testing.B) {
	SetHas(b, 900)
}

func BenchmarkSetAdd500(b *testing.B) {
	SetAdd(b, 500)
}

func BenchmarkSetHas500(b *testing.B) {
	SetHas(b, 500)
}

func BenchmarkSetAdd900(b *testing.B) {
	SetAdd(b, 900)
}

func BenchmarkSetHas100(b *testing.B) {
	SetHas(b, 100)
}
