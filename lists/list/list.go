package list

import (
	"github.com/hovertank3d/stdfp"
	"github.com/hovertank3d/stdfp/lists"
)

type List[T any] struct {
	op     func(int) T
	size   int
	finite bool
}

func FromSlice[T any](s []T) List[T] {
	return NewFinite(len(s), func(i int) T {
		return s[i]
	})
}

func New[T any](op func(int) T) List[T] {
	return List[T]{
		op:     op,
		finite: false,
	}
}

func NewFinite[T any](size int, op func(int) T) List[T] {
	return List[T]{
		op:     op,
		size:   size,
		finite: true,
	}
}

func (l List[T]) Part(a int, b int) lists.FiniteList[T] {
	a, b = l.bounds(a, b)
	return List[T]{
		finite: true,
		size:   b - a,
		op: func(i int) T {
			return l.op(a + i)
		},
	}
}

func (l List[T]) Slice(a int, b int) []T {
	a, b = l.bounds(a, b)
	size := b - a

	s := make([]T, size)
	stdfp.ForEach[T](l, func(i int, t T) {
		s[i] = t
	})

	return s
}

func (l List[T]) Element(i int) T {
	if i > l.size {
		panic("stdfp/lists: out of bounds")
	}
	return l.op(i)
}

func (l List[T]) Size() int {
	return l.size
}

func (l List[T]) bounds(a int, b int) (int, int) {
	if l.finite == false {
		return a, b
	}

	if a > l.size {
		return 0, 0
	} else if a+b > l.size {
		return a, l.size
	}
	return a, b
}
