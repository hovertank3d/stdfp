package stdfp

type (
	listInfo struct {
		infinite bool
		size     int
	}

	list[T any] func(i *listInfo, s int, e int) []T

	List[T any] interface {
		Element(int) T
		Slice(int, int) []T
		Size() int
		Infinite() bool
		Map(func(T) T) List[T]
	}
)

func NewListSlice[T any](s []T) List[T] {
	return list[T](func(i *listInfo, a int, b int) []T {
		if i != nil {
			i.size = len(s)
			i.infinite = false
			return []T{}
		}
		if a > len(s) {
			return []T{}
		}
		if a+b > len(s) {
			return s[a:]
		}
		return s[a:b]
	})
}

func (l list[T]) Element(i int) T {
	return l(nil, i, i)[0]
}

func (l list[T]) Slice(a int, b int) []T {
	return l(nil, a, b)
}

func (l list[T]) Size() int {
	var i listInfo
	l(&i, 0, 0)
	return i.size
}

func (l list[T]) Infinite() bool {
	var i listInfo
	l(&i, 0, 0)
	return i.infinite
}
