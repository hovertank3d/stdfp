package stdfp

func (l list[T]) Map(m func(T) T) List[T] {
	return list[T](func(i *listInfo, a int, b int) (result []T) {
		if i != nil {
			return l(i, 0, 0)
		}
		for _, e := range l.Slice(a, b) {
			result = append(result, m(e))
		}
		return result
	})
}
