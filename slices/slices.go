package slices

// Map maps the elements of a slice of type A into a new slice of type B.
func Map[A, B any](as []A, f func(A) B) []B {
	if as == nil {
		return nil
	}

	bs := make([]B, len(as))
	for i, n := 0, len(as); i < n; i++ {
		bs[i] = f(as[i])
	}
	return bs
}

// FlatMap flat maps the elements of a slice of type A into a new slice of type B.
func FlatMap[A, B any](as []A, f func(A) []B) []B {
	if as == nil {
		return nil
	}

	bs := make([]B, 0, len(as))
	for _, a := range as {
		bs = append(bs, f(a)...)
	}
	return bs
}

// Filter filters all elements of a slice which satisfy a predicate.
func Filter[A any](as []A, f func(A) bool) []A {
	return FlatMap(as, func(a A) []A {
		if f(a) {
			return []A{a}
		}
		return nil
	})
}

// Fold folds the elements of a slice using the specified associative binary operator.
func Fold[A any](as []A, z A, op func(A, A) A) A { return FoldLeft[A, A](as, z, op) }

// FoldLeft Applies a binary operator to a start value and all elements of a slice, going left to right.
func FoldLeft[A, B any](as []A, z B, op func(b B, a A) B) B {
	acc := z
	for _, a := range as {
		acc = op(acc, a)
	}
	return acc
}

// FoldRight applies a binary operator to a start value and all elements of a slice, going right to left.
func FoldRight[A, B any](as []A, z B, op func(a A, b B) B) B {
	if len(as) == 0 {
		return z
	}
	head := as[0]
	tail := as[1:]
	return op(head, FoldRight(tail, z, op))
}

// Distinct returns NEW slice with a unique set of elements.
// Order is maintained with the first occurrence of an element appearing first.
func Distinct[A comparable](as []A) []A {
	if as == nil {
		return nil
	}

	n := len(as)
	set := make(map[any]struct{}, n)

	return Filter(as, func(a A) bool {
		if _, ok := set[a]; ok {
			return false
		} else {
			set[a] = struct{}{}
			return true
		}
	})
}

// Reversed returns NEW slice with elements in reversed order.
func Reversed[A any](as []A) []A {
	if as == nil {
		return nil
	}

	n := len(as)
	res := make([]A, n)
	for i := 0; i < n; i++ {
		res[i] = as[n-1-i]
	}

	return res
}

// ForAll returns true if all elements in a slice evaluate to true by the predicate, otherwise it returns False.
// If the slice is empty or nil, ForAll() will return true.
func ForAll[A any](as []A, f func(A) bool) bool {
	for _, v := range as {
		if !f(v) {
			return false
		}
	}
	return true
}

// ForAny returns true if any element in a slice evaluates to true by the predicate, otherwise it returns False.
// If the slice is empty or nil, ForAny() will return false.
func ForAny[A any](as []A, f func(A) bool) bool {
	for _, v := range as {
		if f(v) {
			return true
		}
	}
	return false
}

func Zip[A, B any](as []A, bs []B) []struct {
	A A
	B B
} {
	n := len(as)

	if n > len(bs) {
		n = len(bs)
	}

	res := make([]struct {
		A A
		B B
	}, n)

	for i := 0; i < n; i++ {
		res[i] = struct {
			A A
			B B
		}{A: as[i], B: bs[i]}
	}

	return res
}
