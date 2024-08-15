package fp

import (
	"fmt"
	"testing"
)

func SelectEven(i int) bool {
	return i%2 == 0
}

func Square(i int) int {
	return i * i
}

func Sum(i int, j int) int {
	return i + j
}

func TestSlice_Reduce(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	t.Logf("got = %v", Slice[int](data).Filter(SelectEven).Map(Square).Reduce(Sum, 0))
}

func TestTransform(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	list := Transform(data, func(t int) string {
		return fmt.Sprintf("%d", t)
	})
	t.Logf("got = %v", list)
}
