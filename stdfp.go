package stdfp

import (
	"github.com/hovertank3d/stdfp/lists"
	"github.com/hovertank3d/stdfp/lists/list"
)

func ForEach[T any](l lists.FiniteList[T], f func(int, T)) {
	for i := 0; i < l.Size(); i++ {
		f(i, l.Element(i))
	}
}

func Repeat[T any](l lists.FiniteList[T]) lists.List[T] {
	return list.New(func(i int) T {
		return l.Element(i % l.Size())
	})
}

func Map[T any](l lists.List[T], m func(T) T) lists.List[T] {
	return list.New(func(i int) T {
		return m(l.Element(i))
	})
}

func MapF[T any](l lists.FiniteList[T], m func(T) T) lists.FiniteList[T] {
	return list.NewFinite(l.Size(), func(i int) T {
		return m(l.Element(i))
	})
}
