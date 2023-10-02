package slices_test

import (
	"reflect"
	"testing"

	"github.com/tdibacco/fp/slices"
)

func TestFoldLeft(t *testing.T) {
	type args[A comparable, B comparable] struct {
		s  []A
		z  B
		op func(b B, a A) B
	}
	type testCase[A comparable, B comparable] struct {
		name string
		args args[A, B]
		want B
	}
	tests := []testCase[int, int]{{
		name: "nil slice",
		args: args[int, int]{s: nil, z: 0, op: func(b, a int) int { return a + b }},
		want: 0,
	}, {
		name: "empty slice",
		args: args[int, int]{s: []int{}, z: 0, op: func(b, a int) int { return a + b }},
		want: 0,
	}, {
		name: "slice of 1",
		args: args[int, int]{s: []int{1}, z: 0, op: func(b, a int) int { return a + b }},
		want: 1,
	}, {
		name: "slice of 2",
		args: args[int, int]{s: []int{1, 2}, z: 0, op: func(b, a int) int { return a + b }},
		want: 3,
	}, {
		name: "slice of 3",
		args: args[int, int]{s: []int{1, 2, 3}, z: 0, op: func(b, a int) int { return a + b }},
		want: 6,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slices.FoldLeft(tt.args.s, tt.args.z, tt.args.op); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoldLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoldRight(t *testing.T) {
	type args[A comparable, B comparable] struct {
		s  []A
		z  B
		op func(a A, b B) B
	}
	type testCase[A comparable, B comparable] struct {
		name string
		args args[A, B]
		want B
	}
	tests := []testCase[int, int]{{
		name: "nil slice",
		args: args[int, int]{s: nil, z: 0, op: func(a, b int) int { return a + b }},
		want: 0,
	}, {
		name: "empty slice",
		args: args[int, int]{s: []int{}, z: 0, op: func(a, b int) int { return a + b }},
		want: 0,
	}, {
		name: "slice of 1",
		args: args[int, int]{s: []int{1}, z: 0, op: func(a, b int) int { return a + b }},
		want: 1,
	}, {
		name: "slice of 2",
		args: args[int, int]{s: []int{1, 2}, z: 0, op: func(a, b int) int { return a + b }},
		want: 3,
	}, {
		name: "slice of 3",
		args: args[int, int]{s: []int{1, 2, 3}, z: 0, op: func(a, b int) int { return a + b }},
		want: 6,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slices.FoldRight(tt.args.s, tt.args.z, tt.args.op); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoldRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[A any, B any] struct {
		as []A
		f  func(A) B
	}
	type testCase[A any, B any] struct {
		name string
		args args[A, B]
		want []B
	}
	tests := []testCase[int, int]{{
		name: "nil",
		args: args[int, int]{as: nil, f: func(int) int { return 0 }},
		want: nil,
	}, {
		name: "empty",
		args: args[int, int]{as: []int{}, f: func(int) int { return 0 }},
		want: []int{},
	}, {
		name: "simple",
		args: args[int, int]{
			as: []int{1, 2, 3},
			f:  func(a int) int { return a + a },
		},
		want: []int{2, 4, 6},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slices.Map(tt.args.as, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatMap(t *testing.T) {
	type args[A any, B any] struct {
		as []A
		f  func(A) []B
	}
	type testCase[A any, B any] struct {
		name string
		args args[A, B]
		want []B
	}
	tests := []testCase[int, int]{{
		name: "nil",
		args: args[int, int]{as: nil, f: func(int) []int { return []int{0} }},
		want: nil,
	}, {
		name: "empty",
		args: args[int, int]{as: []int{}, f: func(int) []int { return []int{0} }},
		want: []int{},
	}, {
		name: "bigger slice",
		args: args[int, int]{
			as: []int{1, 2, 3},
			f:  func(a int) []int { return []int{a, a} },
		},
		want: []int{1, 1, 2, 2, 3, 3},
	}, {
		name: "smaller slice",
		args: args[int, int]{
			as: []int{1, 2, 3},
			f: func(a int) []int {
				if a%2 != 0 {
					return []int{a}
				}
				return nil
			}},
		want: []int{1, 3},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slices.FlatMap(tt.args.as, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlatMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	t.Run("filter odd ints", func(t *testing.T) {
		as := []int{1, 2, 3}
		f := func(a int) bool { return a%2 != 0 }
		want := []int{1, 3}
		if got := slices.Filter(as, f); !reflect.DeepEqual(got, want) {
			t.Errorf("FlatMap() = %v, want %v", got, want)
		}
	})
}

func TestMean(t *testing.T) {
	t.Run("mean", func(t *testing.T) {
		as := []int{1, 2, 3, 4, 5}
		z := [2]int{0, 0}
		f := func(b [2]int, a int) [2]int {
			return [2]int{b[0] + a, b[1] + 1}
		}

		if want, got := [2]int{15, 5}, slices.FoldLeft(as, z, f); !reflect.DeepEqual(got, want) {
			t.Errorf("FoldLeft = %v, want %v", got, want)
		}
	})
}

func TestDistinct(t *testing.T) {
	type args[A any] struct {
		as []A
	}
	type testCase[A any] struct {
		name string
		args args[A]
		want []A
	}
	tests := []testCase[int]{{
		name: "nil",
		args: args[int]{nil},
		want: nil,
	}, {
		name: "empty",
		args: args[int]{[]int{}},
		want: []int{},
	}, {
		name: "simple",
		args: args[int]{[]int{7, 1, 4, 6, 7, 9, 1, 5, 1, 2}},
		want: []int{7, 1, 4, 6, 9, 5, 2},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slices.Distinct(tt.args.as); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReversed(t *testing.T) {
	type args[A any] struct {
		as []A
	}
	type testCase[A any] struct {
		name string
		args args[A]
		want []A
	}
	tests := []testCase[int]{{
		name: "nil",
		args: args[int]{nil},
		want: nil,
	}, {
		name: "empty",
		args: args[int]{[]int{}},
		want: []int{},
	}, {
		name: "simple",
		args: args[int]{[]int{7, 1, 4, 6, 7, 9, 1, 5, 1, 2}},
		want: []int{2, 1, 5, 1, 9, 7, 6, 4, 1, 7},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slices.Reversed(tt.args.as); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reversed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForAll(t *testing.T) {
	type args[A any] struct {
		as []A
		f  func(A) bool
	}
	type testCase[A any] struct {
		name string
		args args[A]
		want bool
	}
	tests := []testCase[bool]{{
		name: "nil",
		args: args[bool]{
			as: nil,
			f:  func(a bool) bool { return a },
		},
		want: true,
	}, {
		name: "empty",
		args: args[bool]{
			as: []bool{},
			f:  func(a bool) bool { return a },
		},
		want: true,
	}, {
		name: "contains a false",
		args: args[bool]{
			as: []bool{true, false, true},
			f:  func(a bool) bool { return a },
		},
		want: false,
	}, {
		name: "all true",
		args: args[bool]{
			as: []bool{true, true, true},
			f:  func(a bool) bool { return a },
		},
		want: true,
	}, {
		name: "all false",
		args: args[bool]{
			as: []bool{false, false, false},
			f:  func(a bool) bool { return a },
		},
		want: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slices.ForAll(tt.args.as, tt.args.f); got != tt.want {
				t.Errorf("ForAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForAny(t *testing.T) {
	type args[A any] struct {
		as []A
		f  func(A) bool
	}
	type testCase[A any] struct {
		name string
		args args[A]
		want bool
	}
	tests := []testCase[bool]{{
		name: "nil",
		args: args[bool]{
			as: nil,
			f:  func(a bool) bool { return a },
		},
		want: false,
	}, {
		name: "empty",
		args: args[bool]{
			as: []bool{},
			f:  func(a bool) bool { return a },
		},
		want: false,
	}, {
		name: "contains a false",
		args: args[bool]{
			as: []bool{true, false, true},
			f:  func(a bool) bool { return a },
		},
		want: true,
	}, {
		name: "all true",
		args: args[bool]{
			as: []bool{true, true, true},
			f:  func(a bool) bool { return a },
		},
		want: true,
	}, {
		name: "all false",
		args: args[bool]{
			as: []bool{false, false, false},
			f:  func(a bool) bool { return a },
		},
		want: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slices.ForAny(tt.args.as, tt.args.f); got != tt.want {
				t.Errorf("ForAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZip(t *testing.T) {
	as := []string{"Trent", "Naoko", "River", "Meadow"}
	bs := []int{49, 50, 16, 14}

	type args[A any, B any] struct {
		as []A
		bs []B
	}
	type testCase[A any, B any] struct {
		name string
		args args[A, B]
		want []struct {
			A A
			B B
		}
	}
	tests := []testCase[string, int]{{
		name: "nil a",
		args: args[string, int]{nil, bs},
		want: []struct {
			A string
			B int
		}{ /* empty */ },
	}, {
		name: "nil b",
		args: args[string, int]{as, nil},
		want: []struct {
			A string
			B int
		}{ /* empty */ },
	}, {
		name: "nil a and b",
		args: args[string, int]{nil, nil},
		want: []struct {
			A string
			B int
		}{ /* empty */ },
	},
		{
			name: "a longer than b",
			args: args[string, int]{as, bs[0 : len(bs)-1]},
			want: []struct {
				A string
				B int
			}{
				{"Trent", 49},
				{"Naoko", 50},
				{"River", 16},
			},
		},
		{
			name: "b longer than a",
			args: args[string, int]{as[0 : len(as)-1], bs},
			want: []struct {
				A string
				B int
			}{
				{"Trent", 49},
				{"Naoko", 50},
				{"River", 16},
			},
		}, {
			name: "names and ages",
			args: args[string, int]{as, bs},
			want: []struct {
				A string
				B int
			}{
				{"Trent", 49},
				{"Naoko", 50},
				{"River", 16},
				{"Meadow", 14},
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slices.Zip(tt.args.as, tt.args.bs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zip() = %v, want %v", got, tt.want)
			}
		})
	}
}
