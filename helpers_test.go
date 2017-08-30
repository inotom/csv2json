package main

import "testing"

func TestString2Int(t *testing.T) {
	var tests = []struct {
		s    string
		want int64
	}{
		{"0", 0},
		{"1", 1},
		{"110", 110},
		{" 110 ", 110},
		{"-1", -1},
		{"-110", -110},
		{" -110 ", -110},
	}

	for _, test := range tests {
		got, _ := string2Int(test.s)
		if got != test.want {
			t.Errorf("string2Int(%s) returned %d, want %d", test.s, got, test.want)
		}
	}
}

func TestString2Float(t *testing.T) {
	var tests = []struct {
		s    string
		want float64
	}{
		{"0.0", 0.0},
		{"1.0", 1.0},
		{"1.01", 1.01},
		{"110.0", 110.0},
		{" 110.0 ", 110.0},
		{"-1.0", -1.0},
		{"-1.01", -1.01},
		{"-110.0", -110.0},
		{" -110.0 ", -110.0},
	}

	for _, test := range tests {
		got, _ := string2Float(test.s)
		if got != test.want {
			t.Errorf("string2Float(%s) returned %f, want %f", test.s, got, test.want)
		}
	}
}
