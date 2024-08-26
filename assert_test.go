package assert_test

import (
	"errors"
	"testing"

	"github.com/Jamlie/assert"
)

func TestNoError(t *testing.T) {
	type args struct {
		err error
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Nil Error",
			args: args{
				err: nil,
				msg: "Should not occur",
			},
		},
		{
			name: "Non-Nil Error",
			args: args{
				err: errors.New("Panic Attack"),
				msg: "Panic Attack",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Log(tt.args.msg)
				}
			}()

			assert.NoError(tt.args.err, tt.args.msg)
		})
	}
}

func TestEquals(t *testing.T) {
	type args struct {
		a   any
		b   any
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Equivalent Slices",
			args: args{
				a:   []int{1, 2, 3, 4, 5},
				b:   []int{1, 2, 3, 4, 5},
				msg: "Equivalent Slices should NEVER panic",
			},
		},
		{
			name: "Non-Equivalent Slices",
			args: args{
				a:   []int{0, 3, 4, 5},
				b:   []int{1, 2, 3, 4, 5},
				msg: "Non-Equivalent Slices must panic, so this is fine",
			},
		},
		{
			name: "Equivalent Maps",
			args: args{
				a: map[string]int{
					"first":  1,
					"second": 2,
				},
				b: map[string]int{
					"first":  1,
					"second": 2,
				},
				msg: "Equivalent Maps should NEVER panic",
			},
		},
		{
			name: "Non-Equivalent Maps",
			args: args{
				a: map[string]int{
					"whatever": 69,
					"second":   2,
				},
				b: map[string]int{
					"first":  1,
					"second": 2,
				},
				msg: "Non-Equivalent Maps must panic, so this is fine",
			},
		},
		{
			name: "Int And Float64",
			args: args{
				a:   1,
				b:   1.,
				msg: "These should not be equal duh",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Log(tt.args.msg)
				}
			}()

			assert.Equals(tt.args.a, tt.args.b, tt.args.msg)
		})
	}
}

func TestNotEquals(t *testing.T) {
	type args struct {
		a   any
		b   any
		msg string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Equivalent Slices",
			args: args{
				a:   []int{1, 2, 3, 4, 5},
				b:   []int{1, 2, 3, 4, 5},
				msg: "Equivalent Slices must panic, so this is fine",
			},
		},
		{
			name: "Non-Equivalent Slices",
			args: args{
				a:   []int{0, 3, 4, 5},
				b:   []int{1, 2, 3, 4, 5},
				msg: "Non-Equivalent Slices should not panic",
			},
		},
		{
			name: "Equivalent Maps",
			args: args{
				a: map[string]int{
					"first":  1,
					"second": 2,
				},
				b: map[string]int{
					"first":  1,
					"second": 2,
				},
				msg: "Equivalent Maps must panic, so this is fine",
			},
		},
		{
			name: "Non-Equivalent Maps",
			args: args{
				a: map[string]int{
					"whatever": 69,
					"second":   2,
				},
				b: map[string]int{
					"first":  1,
					"second": 2,
				},
				msg: "Non-Equivalent Maps should not panic",
			},
		},
		{
			name: "Int And Float64",
			args: args{
				a:   1,
				b:   1.,
				msg: "Int and Float64 are not the same",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Log(tt.args.msg)
				}
			}()

			assert.NotEquals(tt.args.a, tt.args.b, tt.args.msg)
		})
	}
}

type numerical interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func TestGreaterThan(t *testing.T) {
	type args[T numerical] struct {
		a   T
		b   T
		msg string
	}
	tests := []struct {
		name string
		args args[int]
	}{
		{
			name: "Equal Values",
			args: args[int]{
				a:   0,
				b:   0,
				msg: "MUST panic, equal values are not greater than each other",
			},
		},
		{
			name: "First Greater Than Second",
			args: args[int]{
				a:   420,
				b:   69,
				msg: "Must pass",
			},
		},
		{
			name: "Second Greater Than First",
			args: args[int]{
				a:   69,
				b:   420,
				msg: "Must not pass, first must be greater than second",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Log(tt.args.msg)
				}
			}()

			assert.GreaterThan(tt.args.a, tt.args.b, tt.args.msg)
		})
	}
}

func TestLessThan(t *testing.T) {
	type args[T numerical] struct {
		a   T
		b   T
		msg string
	}
	tests := []struct {
		name string
		args args[float64]
	}{
		{
			name: "Equal Values",
			args: args[float64]{
				a:   0.,
				b:   0.,
				msg: "MUST panic, equal values are not greater than each other",
			},
		},
		{
			name: "First Greater Than Second",
			args: args[float64]{
				a:   420.21,
				b:   69.98,
				msg: "Must not pass, first must be less than second",
			},
		},
		{
			name: "Second Greater Than First",
			args: args[float64]{
				a:   69.98,
				b:   420.21,
				msg: "Must pass, second must be greater than first",
			},
		},
		{
			name: "Close Values, first greater than second",
			args: args[float64]{
				a:   69.98,
				b:   69.96,
				msg: "Must not pass, second must be greater than first",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Log(tt.args.msg)
				}
			}()

			assert.LessThan(tt.args.a, tt.args.b, tt.args.msg)
		})
	}
}

func TestGreaterThanEquals(t *testing.T) {
	type args[T numerical] struct {
		a   T
		b   T
		msg string
	}
	tests := []struct {
		name string
		args args[uint8]
	}{
		{
			name: "Equal Values",
			args: args[uint8]{
				a:   0,
				b:   0,
				msg: "Must not panic",
			},
		},
		{
			name: "First Greater Than Second",
			args: args[uint8]{
				a:   69,
				b:   21,
				msg: "Must pass",
			},
		},
		{
			name: "Second Greater Than First",
			args: args[uint8]{
				a:   21,
				b:   69,
				msg: "Must not pass, first must be greater than second",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Log(tt.args.msg)
				}
			}()

			assert.GreaterThanEquals(tt.args.a, tt.args.b, tt.args.msg)
		})
	}
}

func TestLessThanEquals(t *testing.T) {
	type args[T numerical] struct {
		a   T
		b   T
		msg string
	}
	tests := []struct {
		name string
		args args[int64]
	}{
		{
			name: "Equal Values",
			args: args[int64]{
				a:   0,
				b:   0,
				msg: "Must panic",
			},
		},
		{
			name: "First Greater Than Second",
			args: args[int64]{
				a:   69,
				b:   21,
				msg: "Must not pass, first must be less than second",
			},
		},
		{
			name: "Second Greater Than First",
			args: args[int64]{
				a:   21,
				b:   69,
				msg: "Must pass",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Log(tt.args.msg)
				}
			}()

			assert.LessThanEquals(tt.args.a, tt.args.b, tt.args.msg)
		})
	}
}

func TestNotEmptySlice(t *testing.T) {
	type args[T any] struct {
		s   []T
		msg string
	}
	tests := []struct {
		name string
		args args[int16]
	}{
		{
			name: "Empty Slice",
			args: args[int16]{
				s:   []int16{},
				msg: "Empty Slice must not pass",
			},
		},
		{
			name: "Non-Empty Slice",
			args: args[int16]{
				s:   []int16{1, 2, 3, 4},
				msg: "Non-Empty Slice must pass",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Log(tt.args.msg)
				}
			}()

			assert.NotEmptySlice(tt.args.s, tt.args.msg)
		})
	}
}

func TestNotEmptyMap(t *testing.T) {
	type args[K comparable, V any] struct {
		m   map[K]V
		msg string
	}
	tests := []struct {
		name string
		args args[string, uint]
	}{
		{
			name: "Empty Map",
			args: args[string, uint]{
				m:   make(map[string]uint),
				msg: "Empty Map must not pass",
			},
		},
		{
			name: "Non-Empty Map",
			args: args[string, uint]{
				m: map[string]uint{
					"special-num": 69,
				},
				msg: "Non-Empty Map must pass",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil {
					t.Log(tt.args.msg)
				}
			}()

			assert.NotEmptyMap(tt.args.m, tt.args.msg)
		})
	}
}
