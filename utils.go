package usvc

import "golang.org/x/exp/constraints"

// MinInt returns the minimum value of two ints. Use math.Min for floats
func MinInt[T constraints.Integer](x, y T) T {
	if x < y {
		return x
	}
	return y
}

// MaxInt returns the minimum value of two ints. Use math.Max for floats
func MaxInt[T constraints.Integer](x, y T) T {
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
