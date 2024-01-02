package f

// Ternary operator
func If[T any](cond bool, a T, b T) T {
	if cond {
		return a
	} else {
		return b
	}
}

func Me(err error) {
	if err != nil {
		panic(err)
	}
}

func F[A any](err error) {
	if err != nil {
		panic(err)
	}
}

func F1[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}

func F2[A any, B any](a A, b B, err error) (A, B) {
	if err != nil {
		panic(err)
	}
	return a, b
}
