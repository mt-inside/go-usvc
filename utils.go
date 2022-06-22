package usvc

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Ternary[T any](test bool, yes, no T) T {
	if test {
		return yes
	}
	return no
}
