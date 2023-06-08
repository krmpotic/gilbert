package main

import (
	"fmt"
)

func main() {
	a := &Matrix{
		{1, 2},
		{2, 2},
		{5, 7},
	}

	b := &Matrix{
		{0, 1},
		{0, 2},
		{0, 3},
	}

	c, _ := a.Add(b)

	fmt.Println("A: ")
	fmt.Println(a)
	fmt.Println("B: ")
	fmt.Println(b)
	fmt.Println("A+B: ")
	fmt.Println(c)
	fmt.Println("C columns:")
	fmt.Println(c.Cols())
	fmt.Println("AT:")
	fmt.Println(a.T())
	fmt.Println("A * AT:")
	fmt.Println(a.Mul(a.T()))
}
