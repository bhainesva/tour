package main

import (
	"fmt"
	"math"
	"strings"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i:=0;;i++ {
		num := math.Pow(z,2) - x
		den := 2 * z
		n := z - (num/den)
		if n == z {
			return n
		}
		z = n
	}

	return z
}

func CubeRt(x complex128, z complex128) complex128 {
	for i:=0;;i++ {
		num := z * z * z - x
		den := 3 * z * z

		n := z - num/den
		fmt.Println(n)
		if n == z {
			return n
		}
		z = n
	}
	return z
}

func WordCount(s string) map[string]int {
	counts := map[string]int{}
	words := strings.Fields(s)
	for _, word := range words {
		if word == "" { continue }
		counts[word]++
	}
	return counts
}

func Pic(dx, dy int) [][]uint8 {
	out := make([][]uint8, dy)
	for i:=0;i<dy;i++ {
		out[i] = make([]uint8, dx)

		for j:=0;j<dx;j++ {
			//out[i][j] = uint8((i + j) / 2)
			//out[i][j] = uint8(math.Pow(float64(i), float64(j)))
			out[i][j] = uint8(i * j)
		}
	}

	return out
}
