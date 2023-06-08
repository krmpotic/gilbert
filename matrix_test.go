package main

import (
	"reflect"
	"testing"
)

var matrix1 = &Matrix{{2, -2, 1}, {5, 5, 5}}

func TestSize(t *testing.T) {
	in := []*Matrix{
		{{2, 0}, {1, 1}, {3, 5}},
		NewMatrix(7, 17),
		NewIdentityMatrix(12),
		{{}, {}, {}},
		matrix1,
	}
	want := [][]int{
		{3, 2},
		{7, 17},
		{12, 12},
		{0, 0},
		{2, 3},
	}

	for i, matrix := range in {
		if m, n := matrix.Size(); m != want[i][0] || n != want[i][1] {
			t.Fatalf("(*Matrix).Size(): want %dx%d; got %dx%d\n", want[i][0], want[i][1], m, n)
		}
	}
}

func TestRows(t *testing.T) {
	rows := matrix1.Rows()
	if len(rows) != 2 {
		t.Fatalf("len((*Matrix).Rows()): want 2, got %d", len(rows))
	}
	if !reflect.DeepEqual(rows[1], Vector{5, 5, 5}) {
		t.Fatalf("(*Matrix).Rows()[1]: want {5,5,5}; got %v", rows[1])
	}
}

func TestCols(t *testing.T) {
	cols := matrix1.Cols()
	if len(cols) != 3 {
		t.Fatalf("len((*Matrix).Cols()): want 3, got %d", len(cols))
	}
	if !reflect.DeepEqual(cols[1], Vector{-2, 5}) {
		t.Fatalf("(*Matrix).Cols()[1]: want {-2,5}; got %v", cols[1])
	}
}
