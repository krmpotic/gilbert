package main

import (
	"math"
)

type Vector []float64

func (v *Vector) Dot(w *Vector) float64 {
	s := 0.0
	for i := 0; i < len(*v); i++ {
		s += (*v)[i] * (*w)[i]
	}
	return s
}

func (v *Vector) Norm() float64 {
	return math.Sqrt(v.Dot(v))
}
