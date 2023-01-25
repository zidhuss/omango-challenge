package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {

	tests := []struct {
		input int
		want  bool
	}{
		{input: 22222, want: true},
		{input: 75, want: false},
		{input: 99, want: true},
		{input: 100, want: false},
		{input: 4879, want: true},
		{input: 1, want: true},
		{input: 9, want: true},
		{input: 703, want: true},
		{input: 2728, want: true},
	}

	for _, tc := range tests {
		t.Run(strconv.Itoa(tc.input), func(t *testing.T) {
			got := Solve(tc.input)
			if got != tc.want {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		},
		)
	}

}
