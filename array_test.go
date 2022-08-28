package array

import (
	"reflect"
	"strconv"
	"testing"

	"golang.org/x/exp/slices"
)

func TestMap(t *testing.T) {
	tt := map[int]struct {
		input    []string
		callback func(string) int
		want     []int
	}{
		1: {
			input:    []string{"hello", "my", "dear", "friend"},
			callback: func(s string) int { return len(s) },
			want:     []int{5, 2, 4, 6},
		},
		2: {
			input:    []string{"", "m", "de", "end"},
			callback: func(s string) int { return len(s) },
			want:     []int{0, 1, 2, 3},
		},
		3: {
			input:    []string{},
			callback: func(s string) int { return len(s) },
			want:     []int{},
		},
		4: {
			input:    nil,
			callback: func(s string) int { return len(s) },
			want:     []int{},
		},
	}

	for index, tc := range tt {
		got := Map(tc.input, tc.callback)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

func TestFilter(t *testing.T) {
	tt := map[int]struct {
		input    []string
		callback func(string) bool
		want     []string
	}{
		1: {
			input:    []string{"hello", "my", "dear", "friend"},
			callback: func(s string) bool { return len(s) > 4 },
			want:     []string{"hello", "friend"},
		},
		2: {
			input:    []string{"hello", "my", "dear", "friend"},
			callback: func(s string) bool { return len(s) > 8 },
			want:     []string{},
		},
		3: {
			input:    nil,
			callback: func(s string) bool { return len(s) > 8 },
			want:     []string{},
		},
	}

	for index, tc := range tt {
		got := Filter(tc.input, tc.callback)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

func TestFilterErrors(t *testing.T) {
	tt := map[int]struct {
		input    []string
		callback func(string) (int, error)
		want     []int
	}{
		1: {
			input:    []string{"123", "12qwerty3", "123qwerty", "qwerty123"},
			callback: strconv.Atoi,
			want:     []int{123},
		},
		2: {
			input:    []string{"qwerty", "123e"},
			callback: strconv.Atoi,
			want:     []int{},
		},
		3: {
			input:    []string{"5487.2", "123", "", "0", "23"},
			callback: strconv.Atoi,
			want:     []int{123, 0, 23},
		},
	}

	for index, tc := range tt {
		got := FilterErrors(tc.input, tc.callback)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

func TestEvery(t *testing.T) {
	tt := map[int]struct {
		input    []string
		callback func(string) bool
		want     bool
	}{
		1: {
			input:    []string{"hello", "my", "dear", "friend"},
			callback: func(s string) bool { return len(s) >= 4 },
			want:     false,
		},
		2: {
			input:    []string{"hello", "my", "dear", "friend"},
			callback: func(s string) bool { return len(s) >= 2 },
			want:     true,
		},
		3: {
			input:    nil,
			callback: func(s string) bool { return len(s) >= 2 },
			want:     false,
		},
	}

	for index, tc := range tt {
		got := Every(tc.input, tc.callback)
		if got != tc.want {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

func TestSome(t *testing.T) {
	tt := map[int]struct {
		input    []string
		callback func(string) bool
		want     bool
	}{
		1: {
			input:    []string{"hello", "my", "dear", "friend"},
			callback: func(s string) bool { return len(s) == 7 },
			want:     false,
		},
		2: {
			input:    []string{"hello", "my", "dear", "friend"},
			callback: func(s string) bool { return len(s) == 2 },
			want:     true,
		},
		3: {
			input:    nil,
			callback: func(s string) bool { return len(s) == 2 },
			want:     false,
		},
	}

	for index, tc := range tt {
		got := Some(tc.input, tc.callback)
		if got != tc.want {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

func TestIndex(t *testing.T) {
	tt := map[int]struct {
		input []string
		key   string
		want  int
	}{
		1: {
			input: []string{"hello", "my", "dear", "friend"},
			key:   "",
			want:  -1,
		},
		2: {
			input: []string{"hello", "my", "dear", "friend"},
			key:   "my",
			want:  1,
		},
		3: {
			input: nil,
			key:   "my",
			want:  -1,
		},
	}

	for index, tc := range tt {
		got := Index(tc.input, tc.key)
		if got != tc.want {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

func TestIndexFunc(t *testing.T) {
	tt := map[int]struct {
		input    []string
		callback func(string) bool
		want     int
	}{
		1: {
			input:    []string{"hello", "my", "dear", "friend"},
			callback: func(s string) bool { return len(s) == 7 },
			want:     -1,
		},
		2: {
			input:    []string{"hello", "my", "dear", "friend"},
			callback: func(s string) bool { return len(s) == 4 },
			want:     2,
		},
		3: {
			input:    nil,
			callback: func(s string) bool { return len(s) == 7 },
			want:     -1,
		},
	}

	for index, tc := range tt {
		got := IndexFunc(tc.input, tc.callback)
		if got != tc.want {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

func TestContains(t *testing.T) {
	tt := map[int]struct {
		input []string
		key   string
		want  bool
	}{
		1: {
			input: []string{"hello", "my", "dear", "friend"},
			key:   "",
			want:  false,
		},
		2: {
			input: []string{"hello", "my", "dear", "friend"},
			key:   "my",
			want:  true,
		},
		3: {
			input: nil,
			key:   "my",
			want:  false,
		},
	}

	for index, tc := range tt {
		got := Contains(tc.input, tc.key)
		if got != tc.want {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

func TestReduce(t *testing.T) {
	tt := map[int]struct {
		input    []string
		initial  int
		callback func(int, string) int
		want     int
	}{
		1: {
			input:    []string{"hello", "my", "dear", "friend"},
			initial:  0,
			callback: func(sum int, s string) int { return sum + len(s) },
			want:     17,
		},
		2: {
			input:    []string{"hello", "my", "dear", "friend"},
			initial:  10,
			callback: func(sum int, s string) int { return sum + 1 },
			want:     14,
		},
		3: {
			input:    nil,
			initial:  0,
			callback: func(sum int, s string) int { return sum + len(s) },
			want:     0,
		},
	}

	for index, tc := range tt {
		got := Reduce(tc.input, tc.initial, tc.callback)
		if got != tc.want {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

func TestFind(t *testing.T) {
	tt := map[int]struct {
		input    []string
		callback func(string) bool
		want     string
		isFound  bool
	}{
		1: {
			input:    []string{"hello", "my", "dear", "friend"},
			callback: func(s string) bool { return len(s) == 2 },
			want:     "my",
			isFound:  true,
		},
		2: {
			input:    []string{"hello", "my", "dear", "friend"},
			callback: func(s string) bool { return len(s) == 7 },
			want:     "",
			isFound:  false,
		},
		3: {
			input:    nil,
			callback: func(s string) bool { return len(s) == 4 },
			want:     "",
			isFound:  false,
		},
	}

	for index, tc := range tt {
		got, isFound := Find(tc.input, tc.callback)
		if got != tc.want && isFound != tc.isFound {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

func TestFill(t *testing.T) {
	tt := map[int]struct {
		input []string
		value string
		start int
		end   int
		want  []string
	}{
		1: {
			input: []string{"hello", "my", "dear", "friend"},
			value: "world",
			start: 0,
			end:   3,
			want:  []string{"world", "world", "world", "world"},
		},
		2: {
			input: []string{"hello", "my", "dear", "friend"},
			value: "world",
			start: -5,
			end:   7,
			want:  []string{"world", "world", "world", "world"},
		},
		3: {
			input: nil,
			value: "world",
			start: 0,
			end:   3,
			want:  nil,
		},
	}

	for index, tc := range tt {
		got := Fill(tc.input, tc.value, tc.start, tc.end)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

func TestReverse(t *testing.T) {
	tt := map[int]struct {
		input []string
		want  []string
	}{
		1: {
			input: []string{"hello", "my", "dear", "friend"},
			want:  []string{"friend", "dear", "my", "hello"},
		},
		2: {
			input: []string{"hello"},
			want:  []string{"hello"},
		},
		3: {
			input: []string{},
			want:  []string{},
		},
		4: {
			input: nil,
			want:  nil,
		},
	}

	for index, tc := range tt {
		got := Reverse(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

func TestMapKeys(t *testing.T) {
	tt := map[int]struct {
		input map[string]int
		want  []string
	}{
		1: {
			input: map[string]int{"hello": 1, "my": 2, "dear": 3, "friend": 4},
			want:  []string{"dear", "friend", "hello", "my"},
		},
		2: {
			input: map[string]int{"hello": 1},
			want:  []string{"hello"},
		},
		3: {
			input: map[string]int{},
			want:  []string{},
		},
		4: {
			input: nil,
			want:  []string{},
		},
	}
	for index, tc := range tt {
		got := MapKeys(tc.input)
		slices.Sort(got)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

func TestMapValues(t *testing.T) {
	tt := map[int]struct {
		input map[string]int
		want  []int
	}{
		1: {
			input: map[string]int{"hello": 1, "my": 2, "dear": 3, "friend": 4},
			want:  []int{1, 2, 3, 4},
		},
		2: {
			input: map[string]int{"hello": 1},
			want:  []int{1},
		},
		3: {
			input: map[string]int{},
			want:  []int{},
		},
		4: {
			input: nil,
			want:  []int{},
		},
	}
	for index, tc := range tt {
		got := MapValues(tc.input)
		slices.Sort(got)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

func TestForEach(t *testing.T) {
	tt := map[int]struct {
		input []string
		want  []int
	}{
		1: {
			input: []string{"hello", "my", "dear", "friend"},
			want:  []int{5, 2, 4, 6},
		},
		2: {
			input: []string{"hello"},
			want:  []int{5},
		},
		3: {
			input: []string{},
			want:  []int{},
		},
		4: {
			input: nil,
			want:  []int{},
		},
	}

	for index, tc := range tt {
		got := []int{}
		ForEach(tc.input, func(s string) { got = append(got, len(s)) })
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}

func TestRange(t *testing.T) {
	tt := map[int]struct {
		input []string
		start int
		n     int
		want  []string
	}{
		1: {
			input: []string{"hello", "my", "dear", "friend"},
			start: 1,
			n:     -100,
			want:  []string{"my", "dear", "friend"},
		},
		2: {
			input: []string{"hello", "my", "dear", "friend"},
			start: 1,
			n:     0,
			want:  []string{"my", "dear", "friend"},
		},
		3: {
			input: []string{"hello", "my", "dear", "friend"},
			start: 1,
			n:     100,
			want:  []string{"my", "dear", "friend"},
		},
		4: {
			input: []string{"hello", "my", "dear", "friend"},
			start: -100,
			n:     1,
			want:  []string{"hello"},
		},
		5: {
			input: []string{"hello", "my", "dear", "friend"},
			start: 0,
			n:     1,
			want:  []string{"hello"},
		},
		6: {
			input: []string{"hello", "my", "dear", "friend"},
			start: 100,
			n:     1,
			want:  []string{},
		},
		7: {
			input: []string{"hello", "my", "dear", "friend"},
			start: -100,
			n:     -100,
			want:  []string{"hello", "my", "dear", "friend"},
		},
		8: {
			input: []string{"hello", "my", "dear", "friend"},
			start: 0,
			n:     0,
			want:  []string{"hello", "my", "dear", "friend"},
		},
		9: {
			input: []string{"hello", "my", "dear", "friend"},
			start: 100,
			n:     100,
			want:  []string{},
		},
		10: {
			input: []string{"hello", "my", "dear", "friend"},
			start: 1,
			n:     2,
			want:  []string{"my", "dear"},
		},
		11: {
			input: []string{"hello", "my", "dear", "friend"},
			start: 0,
			n:     4,
			want:  []string{"hello", "my", "dear", "friend"},
		},
	}

	for index, tc := range tt {
		got := Range(tc.input, tc.start, tc.n)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Test #%d: got = %v; want = %v", index, got, tc.want)
		}
	}
}
