package main

import (
	"testing"
)

func TestT(t *testing.T) {
	in := []*Matrix{
		{{2, 3}, {0, 1}},
		{{1, 2, 3, 4}},
	}
	want := []*Matrix{
		{{2, 0}, {3, 1}},
		{{1}, {2}, {3}, {4}},
	}

	for i, m := range in {
		if !m.T().Equal(want[i]) {
			t.Errorf("T(ranpose): failed test %d.", i)
		}
	}
}
