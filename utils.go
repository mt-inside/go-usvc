package usvc

func Ternary[T any](test bool, yes, no T) T {
	if test {
		return yes
	}
	return no
}
