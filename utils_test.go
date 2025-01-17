package utils

import (
	"reflect"
	"testing"
)

func fnTrue(_ string) bool {
	return true
}

func fnLength2(s string) bool {
	return len(s) == 2
}

func TestAny(t *testing.T) {
	type args[T any] struct {
		list []T
		fn   func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[string]{
		{
			name: "empty",
			args: struct {
				list []string
				fn   func(string) bool
			}{list: make([]string, 0), fn: fnTrue},
			want: false,
		},
		{
			name: "not empty",
			args: struct {
				list []string
				fn   func(string) bool
			}{list: []string{"xd"}, fn: fnTrue},
			want: true,
		},
		{
			name: "not length 2",
			args: struct {
				list []string
				fn   func(string) bool
			}{list: []string{"xdd"}, fn: fnLength2},
			want: false,
		},
		{
			name: "length 2",
			args: struct {
				list []string
				fn   func(string) bool
			}{list: []string{"xdd", "xd"}, fn: fnLength2},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Any(tt.args.list, tt.args.fn); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

type Obj struct {
	Id   int
	Name string
}

func TestFilter(t *testing.T) {
	type args[T any] struct {
		list []T
		fn   func(T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[Obj]{
		{
			name: "1",
			args: struct {
				list []Obj
				fn   func(Obj) bool
			}{list: []Obj{
				{Id: 1, Name: "a"},
				{Id: 2, Name: "b"},
				{Id: 3, Name: "c"},
				{Id: 4, Name: "d"},
			}, fn: func(obj Obj) bool {
				return obj.Id == 2 || obj.Name == "c"
			}},
			want: []Obj{
				{Id: 2, Name: "b"},
				{Id: 3, Name: "c"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.list, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterNotNil(t *testing.T) {
	type args[T any] struct {
		list []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[*Obj]{
		{
			name: "1",
			args: struct {
				list []*Obj
			}{list: []*Obj{
				nil,
				{Id: 1, Name: "a"},
				{Id: 2, Name: "b"},
				nil,
				{Id: 3, Name: "c"},
				{Id: 4, Name: "d"},
				nil,
			}},
			want: []*Obj{
				{Id: 1, Name: "a"},
				{Id: 2, Name: "b"},
				{Id: 3, Name: "c"},
				{Id: 4, Name: "d"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterNotNil(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[T any, R any] struct {
		list []T
		fn   func(T) R
	}
	type testCase[T any, R any] struct {
		name string
		args args[T, R]
		want []R
	}
	tests := []testCase[Obj, string]{
		{
			name: "1",
			args: struct {
				list []Obj
				fn   func(Obj) string
			}{list: []Obj{
				{Id: 1, Name: "a"},
				{Id: 2, Name: "b"},
				{Id: 3, Name: "c"},
				{Id: 4, Name: "d"},
			}, fn: func(obj Obj) string {
				return obj.Name
			}},
			want: []string{"a", "b", "c", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.list, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMap(t *testing.T) {
	type args[T any, K comparable] struct {
		list []T
		fn   func(T) K
	}
	type testCase[T any, K comparable] struct {
		name string
		args args[T, K]
		want map[K][]T
	}
	tests := []testCase[string, int]{
		{
			name: "length map",
			args: args[string, int]{
				list: []string{"xd", "xdd", "xddd", "aa"},
				fn: func(s string) int {
					return len(s)
				},
			},
			want: map[int][]string{
				2: {"xd", "aa"},
				3: {"xdd"},
				4: {"xddd"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMap(tt.args.list, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
