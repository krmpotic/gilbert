package main

func (a *Matrix) Add(b *Matrix) *Matrix {
	m, n := a.Size()
	if mb, nb := b.Size(); !(m == mb && n == nb) {
		panic("add: bad size")
	}

	c := a.Copy()

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			(*c)[i][j] = (*a)[i][j] + (*b)[i][j]
		}
	}

	return c
}

func (a *Matrix) Scale(k float64) *Matrix {
	m, n := a.Size()
	c := a.Copy()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			(*c)[i][j] *= k
		}
	}
	return c
}

func (a *Matrix) Sub(b *Matrix) *Matrix {
	return a.Add(b.Scale(-1.0))
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
		panic("multiply: bad size")
	}
	c := NewMatrix(ma, nb)
	for i, row := range a.Rows() {
		for j, col := range b.Cols() {
			(*c)[i][j] = row.Dot(&col)
		}
	}
	return c
}
