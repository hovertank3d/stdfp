package stdfp

import (
	"github.com/hovertank3d/stdfp/lists"
	"github.com/hovertank3d/stdfp/lists/list"
)

func ForEach[T any](list lists.FiniteList[T], f func(int, T)) {
	for i := 0; i < list.Size(); i++ {
		f(i, list.Element(i))
	}
}

func Map[T any](l lists.List[T], m func(T) T) lists.List[T] {
	return list.New(func(i int) T {
		return m(l.Element(i))
	})
}
