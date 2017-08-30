package main

import "testing"

func TestIsSkip(t *testing.T) {
	var tests = []struct {
		readAll bool
		index   int
		want    bool
	}{
		{true, 0, false},
		{false, 0, true},
		{true, 1, false},
		{false, 1, false},
	}

	for _, test := range tests {
		var f csvFormat
		f.ReadAll = test.readAll
		got := f.isSkip(test.index)
		if got != test.want {
			t.Errorf("f.isSkip(%d) returned %v, want %v", test.index, got, test.want)
		}
	}
}
