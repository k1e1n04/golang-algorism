package main

import (
	"fmt"
	"github.com/k1e1n04/golang-algorism/utils"
	"math"
	"slices"
)

func main() {
	var N, K int
	fmt.Scan(&N)
	fmt.Scan(&K)
	a := make([]int, N)
	b := make([]int, N)
	for i := range a {
		fmt.Scan(&a[i])
	}
	for i := range b {
		fmt.Scan(&b[i])
	}

	slices.Sort(b)

	minV := math.MaxInt
	for i := range a {
		idx := utils.LowerBound(b, K-a[i])
		if idx >= len(b) {
			continue
		}
		v := b[idx]

		if minV > a[i]+v {
			minV = a[i] + v
		}
	}

	fmt.Println(minV)
}
