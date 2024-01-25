package lists

type List[T any] interface {
	Part(int, int) FiniteList[T]
	Slice(int, int) []T
	Element(int) T
}

type FiniteList[T any] interface {
	List[T]
	// Size returns size of finite list
	// for infinite it should return -1
	Size() int
}
