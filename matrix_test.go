package main

import (
	"testing"
)

func TestSize(t *testing.T) {
	in := []*Matrix{
		{{2, 0}, {1, 1}, {3, 5}},
		NewMatrix(7, 17),
		NewIdentityMatrix(12),
		{{}, {}, {}},
	}
	want := [][]int{
		{3, 2},
		{7, 17},
		{12, 12},
		{0, 0},
	}

	for i, matrix := range in {
		if m, n := matrix.Size(); m != want[i][0] || n != want[i][1] {
			t.Fatalf("Size(): want %dx%d; got %dx%d\n", want[i][0], want[i][1], m, n)
		}
	}
}
