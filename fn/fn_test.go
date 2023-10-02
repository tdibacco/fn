package fn_test

import (
	"reflect"
	"testing"

	"github.com/tdibacco/fp/fn"
)

func TestCurry(t *testing.T) {
	type args[A any, B any, C any] struct {
		op func(a A, b B) C
		a  A
		b  B
	}
	type testCase[A any, B any, C any] struct {
		name string
		args args[A, B, C]
		want C
	}
	tests := []testCase[int, int, int]{{
		name: "inc",
		args: args[int, int, int]{
			op: func(a, b int) (c int) { return a + b },
			a:  1,
			b:  10,
		},
		want: 11,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Curry(tt.args.op)(tt.args.a)(tt.args.b); got != tt.want {
				t.Errorf("Curry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	add := func(a int) int { return a + 1 }
	mul := func(a int) int { return a * 2 }

	// compose add(mul(x))
	if want, got := 5, fn.Compose(add, mul)(2); got != want {
		t.Errorf("Compose() = %v, want %v", got, want)
	}

	sub := func(a int) int { return a - 1 }

	// compose add(mul(sub(x))))
	if want, got := 3, fn.Compose(fn.Compose(add, mul), sub)(2); got != want {
		t.Errorf("Compose() = %v, want %v", got, want)
	}
}

func TestAndThen(t *testing.T) {
	add := func(a int) int { return a + 1 }
	mul := func(a int) int { return a * 2 }

	// andThen = mul(add(x))
	if want, got := 6, fn.AndThen(add, mul)(2); got != want {
		t.Errorf("AndThen() = %v, want %v", got, want)
	}

	sub := func(a int) int { return a - 1 }

	// andThen sub(mul(add(x)))
	if want, got := 5, fn.AndThen(fn.AndThen(add, mul), sub)(2); got != want {
		t.Errorf("AndThen() = %v, want %v", got, want)
	}
}

func TestNot(t *testing.T) {
	passThru := func(b bool) bool {
		return b
	}

	type testCase[A any] struct {
		name string
		args bool
		want bool
	}
	tests := []testCase[bool]{{
		name: "false -> true",
		args: false,
		want: true,
	}, {
		name: "true -> false",
		args: true,
		want: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Not(passThru)(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Not() = %v, want %v", got, tt.want)
			}
		})
	}
}
