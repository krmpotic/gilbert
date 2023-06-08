package main

import (
	"fmt"
)

type Matrix [][]float64

func NewMatrix(m, n int) *Matrix {
	var a Matrix
	a = make([][]float64, m)
	for i := 0; i < m; i++ {
		a[i] = make([]float64, n)
	}
	return &a
}

func NewIdentityMatrix(m int) *Matrix {
	a := NewMatrix(m, m)
	for i := 0; i < m; i++ {
		(*a)[i][i] = 1
	}
	return a
}

func (a *Matrix) String() string {
	m, n := a.Size()
	s := ""
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			s += fmt.Sprintf("%v ", (*a)[i][j])
		}
		s += "\n"
	}
	s += "\n"
	return s
}

func (a *Matrix) Size() (int, int) {
	if a == nil || (*a) == nil || (*a)[0] == nil {
		return 0, 0
	}
	return len(*a), len((*a)[0])
}

func (a *Matrix) Add(b *Matrix) (*Matrix, error) {
	m, n := a.Size()
	if mb, nb := b.Size(); !(m == mb && n == nb) {
		return nil, fmt.Errorf("bad size")
	}

	c := a.Copy()

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			(*c)[i][j] = (*a)[i][j] + (*b)[i][j]
		}
	}

	return c, nil
}

func (a *Matrix) Copy() *Matrix {
	m, n := a.Size()
	c := NewMatrix(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			(*c)[i][j] = (*a)[i][j]
		}
	}
	return c
}

func (a *Matrix) Rows() []Vector {
	m, n := a.Size()
	var v []Vector
	v = make([]Vector, m)
	for i := 0; i < m; i++ {
		v[i] = make(Vector, n)
		for j := 0; j < n; j++ {
			v[i][j] = (*a)[i][j]
		}
	}
	return v
}

func (a *Matrix) Cols() []Vector {
	m, n := a.Size()
	var v []Vector
	v = make([]Vector, n)
	for j := 0; j < n; j++ {
		v[j] = make(Vector, m)
		for i := 0; i < m; i++ {
			v[j][i] = (*a)[i][j]
		}
	}
	return v
}

func (a *Matrix) T() *Matrix {
	m, n := a.Size()
	t := NewMatrix(n, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			(*t)[j][i] = (*a)[i][j]
		}
	}
	return t
}

func (a *Matrix) Mul(b *Matrix) *Matrix {
	ma, na := a.Size()
	mb, nb := b.Size()
	if na != mb {
		// todo: return error
		return nil
	}
	c := NewMatrix(ma, nb)
	for i, row := range a.Rows() {
		for j, col := range b.Cols() {
			(*c)[i][j] = row.Dot(&col)
		}
	}
	return c
}
