package fn

// Identity a function that returns its input value.
func Identity[T any](x T) T {
	return x
}

// Not negates a function boolean results.
func Not[T any](f func(T) bool) func(T) bool {
	return func(v T) bool { return !f(v) }
}

// Curry converts a function 'op' of two arguments into a function of one argument that partially applies 'op'.
func Curry[A, B, C any](op func(a A, b B) C) func(A) func(B) C {
	return func(a A) func(B) C {
		return func(b B) C {
			return op(a, b)
		}
	}
}

// Compose composes two functions so that C = f(g(A)).
func Compose[A, B, C any](f func(b B) C, g func(a A) B) func(A) C {
	return func(a A) C { return f(g(a)) }
}

// AndThen applies function f first and then g.
func AndThen[A, B, C any](f func(b B) C, g func(c C) A) func(B) A {
	return func(b B) A { return g(f(b)) }
}
