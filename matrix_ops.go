package main

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
