package main

import (
	"math/rand"
	"sync"
	"testing"
)

type Set struct {
	sync.RWMutex
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
	s.RLock()
	defer s.RUnlock()
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

func SetWork(b *testing.B, message string, border float32) {
	var set = NewSet()

	b.Run(message, func(b *testing.B) {
		b.SetParallelism(10000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() <= border {
					set.Add(1)
				} else {
					set.Has(1)
				}
			}
		})
	})
}

func BenchmarkSetWork10_90(b *testing.B) {
	SetWork(b, "Mutex: 10% write 90% read", 0.1)
}

func BenchmarkSetWork50_50(b *testing.B) {
	SetWork(b, "Mutex: 50% write 50% read", 0.5)
}

func BenchmarkSetWork90_10(b *testing.B) {
	SetWork(b, "Mutex: 90% write 10% read", 0.9)
}
