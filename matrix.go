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
	if len(*a) == 0 || len((*a)[0]) == 0 {
		return 0, 0
	}
	return len(*a), len((*a)[0])
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
