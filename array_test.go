package array

import (
	"testing"
)

var tt1 = []struct {
	input    []string
	callback func(string) int
	output   []int
}{
	{
		input:    []string{"hello", "my", "dear", "friend"},
		callback: func(s string) int { return len(s) },
		output:   []int{5, 2, 4, 6},
	},
	{
		input:    []string{"", "m", "de", "end"},
		callback: func(s string) int { return len(s) },
		output:   []int{0, 1, 2, 3},
	},
	{
		input:    []string{},
		callback: func(s string) int { return len(s) },
		output:   []int{},
	},
	{
		input:    nil,
		callback: func(s string) int { return len(s) },
		output:   []int{},
	},
}

func TestMap(t *testing.T) {
	for _, tc := range tt1 {
		got := Map(tc.input, tc.callback)
		if got == nil && len(got) == len(tc.input) {
			t.Errorf("got = %v; want = %v", got, tc.output)
		}
	}
}

var tt2 = []struct {
	input    []string
	callback func(string) bool
	output   []string
}{
	{
		input:    []string{"hello", "my", "dear", "friend"},
		callback: func(s string) bool { return len(s) >= 4 },
		output:   []string{"hello", "friend"},
	},
	{
		input:    []string{"hello", "my", "dear", "friend"},
		callback: func(s string) bool { return len(s) > 8 },
		output:   []string{},
	},
	{
		input:    nil,
		callback: func(s string) bool { return len(s) > 8 },
		output:   []string{},
	},
}

func TestFilter(t *testing.T) {
	for _, tc := range tt2 {
		got := Filter(tc.input, tc.callback)
		if got == nil && len(got) != len(tc.output) {
			t.Errorf("got = %v; want = %v", len(got), len(tc.output))
		}
	}
}

var tt3 = []struct {
	input    []string
	callback func(string) bool
	output   bool
}{
	{
		input:    []string{"hello", "my", "dear", "friend"},
		callback: func(s string) bool { return len(s) >= 4 },
		output:   false,
	},
	{
		input:    []string{"hello", "my", "dear", "friend"},
		callback: func(s string) bool { return len(s) >= 2 },
		output:   true,
	},
	{
		input:    nil,
		callback: func(s string) bool { return len(s) >= 2 },
		output:   false,
	},
}

func TestEvery(t *testing.T) {
	for _, tc := range tt3 {
		got := Every(tc.input, tc.callback)
		if got != tc.output {
			t.Errorf("got = %v; want = %v", got, tc.output)
		}
	}
}

var tt4 = []struct {
	input    []string
	callback func(string) bool
	output   bool
}{
	{
		input:    []string{"hello", "my", "dear", "friend"},
		callback: func(s string) bool { return len(s) == 7 },
		output:   false,
	},
	{
		input:    []string{"hello", "my", "dear", "friend"},
		callback: func(s string) bool { return len(s) == 2 },
		output:   true,
	},
	{
		input:    nil,
		callback: func(s string) bool { return len(s) == 2 },
		output:   false,
	},
}

func TestSome(t *testing.T) {
	for _, tc := range tt4 {
		got := Some(tc.input, tc.callback)
		if got != tc.output {
			t.Errorf("got = %v; want = %v", got, tc.output)
		}
	}
}

var tt5 = []struct {
	input  []string
	key    string
	output int
}{
	{
		input:  []string{"hello", "my", "dear", "friend"},
		key:    "",
		output: -1,
	},
	{
		input:  []string{"hello", "my", "dear", "friend"},
		key:    "my",
		output: 1,
	},
	{
		input:  nil,
		key:    "my",
		output: -1,
	},
}

func TestIndex(t *testing.T) {
	for _, tc := range tt5 {
		got := Index(tc.input, tc.key)
		if got != tc.output {
			t.Errorf("got = %v; want = %v", got, tc.output)
		}
	}
}

var tt6 = []struct {
	input    []string
	callback func(string) bool
	output   int
}{
	{
		input:    []string{"hello", "my", "dear", "friend"},
		callback: func(s string) bool { return len(s) == 7 },
		output:   -1,
	},
	{
		input:    []string{"hello", "my", "dear", "friend"},
		callback: func(s string) bool { return len(s) == 4 },
		output:   2,
	},
	{
		input:    nil,
		callback: func(s string) bool { return len(s) == 7 },
		output:   -1,
	},
}

func TestIndexFunc(t *testing.T) {
	for _, tc := range tt6 {
		got := IndexFunc(tc.input, tc.callback)
		if got != tc.output {
			t.Errorf("got = %v; want = %v", got, tc.output)
		}
	}
}

var tt7 = []struct {
	input  []string
	key    string
	output bool
}{
	{
		input:  []string{"hello", "my", "dear", "friend"},
		key:    "",
		output: false,
	},
	{
		input:  []string{"hello", "my", "dear", "friend"},
		key:    "my",
		output: true,
	},
	{
		input:  nil,
		key:    "my",
		output: false,
	},
}

func TestContains(t *testing.T) {
	for _, tc := range tt7 {
		got := Contains(tc.input, tc.key)
		if got != tc.output {
			t.Errorf("got = %v; want = %v", got, tc.output)
		}
	}
}

var tt8 = []struct {
	input    []string
	initial  int
	callback func(int, string) int
	output   int
}{
	{
		input:    []string{"hello", "my", "dear", "friend"},
		initial:  0,
		callback: func(sum int, s string) int { return sum + len(s) },
		output:   17,
	},
	{
		input:    []string{"hello", "my", "dear", "friend"},
		initial:  10,
		callback: func(sum int, s string) int { return sum + 1 },
		output:   14,
	},
	{
		input:    nil,
		initial:  0,
		callback: func(sum int, s string) int { return sum + len(s) },
		output:   0,
	},
}

func TestReduce(t *testing.T) {
	for _, tc := range tt8 {
		got := Reduce(tc.input, tc.initial, tc.callback)
		if got != tc.output {
			t.Errorf("got = %v; want = %v", got, tc.output)
		}
	}
}

var tt9 = []struct {
	input    []string
	callback func(string) bool
	output   string
	isFound  bool
}{
	{
		input:    []string{"hello", "my", "dear", "friend"},
		callback: func(s string) bool { return len(s) == 2 },
		output:   "my",
		isFound:  true,
	},
	{
		input:    []string{"hello", "my", "dear", "friend"},
		callback: func(s string) bool { return len(s) == 7 },
		output:   "",
		isFound:  false,
	},
	{
		input:    nil,
		callback: func(s string) bool { return len(s) == 4 },
		output:   "",
		isFound:  false,
	},
}

func TestFind(t *testing.T) {
	for _, tc := range tt9 {
		got, isFound := Find(tc.input, tc.callback)
		if got != tc.output && isFound != tc.isFound {
			t.Errorf("got = %v; want = %v", got, tc.output)
		}
	}
}

var tt10 = []struct {
	input  []string
	value  string
	start  int
	end    int
	output []string
}{
	{
		input:  []string{"hello", "my", "dear", "friend"},
		value:  "world",
		start:  0,
		end:    3,
		output: []string{"world", "world", "world", "world"},
	},
	{
		input:  []string{"hello", "my", "dear", "friend"},
		value:  "world",
		start:  -5,
		end:    7,
		output: []string{"world", "world", "world", "world"},
	},
	{
		input:  nil,
		value:  "world",
		start:  0,
		end:    3,
		output: []string{},
	},
}

func TestFill(t *testing.T) {
	for _, tc := range tt10 {
		got := Fill(tc.input, tc.value, tc.start, tc.end)
		if len(got) != len(tc.input) {
			t.Errorf("got = %v; want = %v", got, tc.output)
		}
	}
}

var tt11 = []struct {
	input  []string
	output []string
}{
	{
		input:  []string{"hello", "my", "dear", "friend"},
		output: []string{"friend", "dear", "my", "hello"},
	},
	{
		input:  []string{"hello"},
		output: []string{"hello"},
	},
	{
		input:  []string{},
		output: []string{},
	},
	{
		input:  nil,
		output: []string{},
	},
}

func TestReverse(t *testing.T) {
	for _, tc := range tt11 {
		got := Reverse(tc.input)
		if len(got) != len(tc.input) {
			t.Errorf("got = %v; want = %v", got, tc.output)
		}
	}
}

var tt12 = []struct {
	input  map[string]int
	output []string
}{
	{
		input:  map[string]int{"hello": 1, "my": 2, "dear": 3, "friend": 4},
		output: []string{"hello", "my", "dear", "friend"},
	},
	{
		input:  map[string]int{"hello": 1},
		output: []string{"hello"},
	},
	{
		input:  map[string]int{},
		output: []string{},
	},
	{
		input:  nil,
		output: []string{},
	},
}

func TestMapKeys(t *testing.T) {
	for _, tc := range tt12 {
		got := MapKeys(tc.input)
		if got == nil || len(got) != len(tc.input) {
			t.Errorf("got = %v; want = %v", got, tc.output)
		}
	}
}

var tt13 = []struct {
	input  map[string]int
	output []int
}{
	{
		input:  map[string]int{"hello": 1, "my": 2, "dear": 3, "friend": 4},
		output: []int{1, 2, 3, 4},
	},
	{
		input:  map[string]int{"hello": 1},
		output: []int{1},
	},
	{
		input:  map[string]int{},
		output: []int{},
	},
	{
		input:  nil,
		output: []int{},
	},
}

func TestMapValues(t *testing.T) {
	for _, tc := range tt13 {
		got := MapValues(tc.input)
		if got == nil || len(got) != len(tc.input) {
			t.Errorf("got = %v; want = %v", got, tc.output)
		}
	}
}

var tt14 = []struct {
	input  []string
	output []int
}{
	{
		input:  []string{"hello", "my", "dear", "friend"},
		output: []int{5, 2, 4, 6},
	},
	{
		input:  []string{"hello"},
		output: []int{1},
	},
	{
		input:  []string{},
		output: []int{},
	},
	{
		input:  nil,
		output: []int{},
	},
}

func TestForEach(t *testing.T) {
	for _, tc := range tt14 {
		got := []int{}
		ForEach(tc.input, func(s string) { got = append(got, len(s)) })
		if got == nil || len(got) != len(tc.input) {
			t.Errorf("got = %v; want = %v", got, tc.output)
		}
	}
}
