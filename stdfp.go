package stdfp

import "github.com/hovertank3d/stdfp/lists"

func ForEach[T any](list lists.FiniteList[T], f func(int, T)) {
	for i := 0; i < list.Size(); i++ {
		f(i, list.Element(i))
	}
}
