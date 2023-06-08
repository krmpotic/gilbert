package main

import "sort"

func (a *Matrix) Block(r []int, c []int) [][]*Matrix {
	m, n := a.Size()
	r = append(r, 0)
	r = append(r, m)
	sort.Sort(sort.IntSlice(r))
	c = append(c, 0)
	c = append(c, n)
	sort.Sort(sort.IntSlice(c))

	blk := make([][]*Matrix, len(r)-1)
	for i := 0; i < len(r)-1; i++ {
		blk[i] = make([]*Matrix, len(c)-1)
		for j := 0; j < len(c)-1; j++ {
			blk[i][j] = a.Window(r[i], c[j], r[i+1], c[j+1])
		}
	}
	return blk
}

func (a *Matrix) Window(i0, j0, i1, j1 int) *Matrix {
	srt := func(a, b int) (int, int) {
		if a > b {
			return b, a
		}
		return a, b
	}
	i0, i1 = srt(i0, i1)
	j0, j1 = srt(j0, j1)

	w := NewMatrix(i1-i0, j1-j0)

	for i := i0; i < i1; i++ {
		for j := j0; j < j1; j++ {
			(*w)[i][j] = (*a)[i0+i][j0+j]
		}
	}

	return w
}

func Augment(blk [][]*Matrix) (aug *Matrix, r []int, c []int) {
	//todo: add size check
	ma, na := 0, 0
	for i := 0; i < len(blk); i++ {
		m, _ := blk[i][0].Size()
		ma += m
	}
	for j := 0; j < len(blk[0]); j++ {
		_, n := blk[0][j].Size()
		na += n
	}

	aug = NewMatrix(ma, na)
	ins := func(i0, j0 int, a *Matrix) {
		m, n := a.Size()
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				(*aug)[i0+i][j0+j] = (*a)[i][j]
			}
		}
	}

	ma = 0
	for i := 0; i < len(blk); i++ {
		na = 0
		for j := 0; j < len(blk[0]); j++ {
			ins(ma, na, blk[i][j])
			_, n := blk[0][j].Size()
			na += n
			if i == 0 {
				c = append(c, na)
			}
		}
		m, _ := blk[i][0].Size()
		ma += m
		r = append(r, ma)
	}

	r = r[:len(r)-1]
	c = c[:len(c)-1]

	return aug, r, c
}
