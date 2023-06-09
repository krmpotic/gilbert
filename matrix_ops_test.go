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
			t.Errorf("T(ranspose): failed test %d.", i)
		}
	}
}

func TestAdd(t *testing.T) {
	in := [][2]*Matrix{
		{{{2, 3}, {0, 1}}, {{1, 2}, {3, 4}}},
	}
	want := []*Matrix{
		{{3, 5}, {3, 5}},
	}

	for i, m := range in {
		if !m[0].Add(m[1]).Equal(want[i]) {
			t.Errorf("Add: failed test %d.", i)
		}
	}
}

func TestScale(t *testing.T) {
	in := []*Matrix{
		{{2, 3}, {0, 1}},
	}
	want := []*Matrix{
		{{6, 9}, {0, 3}},
	}

	for i, m := range in {
		if !m.Scale(3.0).Equal(want[i]) {
			t.Errorf("Scale: failed test %d.", i)
		}
	}
}

func TestSub(t *testing.T) {
	in := [][2]*Matrix{
		{{{2, 3}, {0, 1}}, {{1, 2}, {3, 4}}},
	}
	want := []*Matrix{
		{{1, 1}, {-3, -3}},
	}

	for i, m := range in {
		if !m[0].Sub(m[1]).Equal(want[i]) {
			t.Errorf("Sub: failed test %d.", i)
		}
	}
}

func TestMul(t *testing.T) {
	in := [][2]*Matrix{
		{{{2, 3}, {0, 1}}, {{1, 2}, {3, 4}}},
	}
	want := []*Matrix{
		{{11, 16}, {3, 4}},
	}

	for i, m := range in {
		if !m[0].Mul(m[1]).Equal(want[i]) {
			t.Errorf("Mul: failed test %d.", i)
		}
	}
}
