package fp

// Slice 是一个泛型切片类型
type Slice[T any] []T

// Filter 函数根据条件过滤切片中的元素
func (s Slice[T]) Filter(condition func(T) bool) Slice[T] {
	var result Slice[T]
	for _, v := range s {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map 函数对切片中的每个元素应用变换函数
func (s Slice[T]) Map(transform func(T) T) Slice[T] {
	var result Slice[T]
	for _, v := range s {
		result = append(result, transform(v))
	}
	return result
}

// Reduce 函数将切片归约为一个值
func (s Slice[T]) Reduce(accumulator func(T, T) T, initial T) T {
	result := initial
	for _, v := range s {
		result = accumulator(result, v)
	}
	return result
}
