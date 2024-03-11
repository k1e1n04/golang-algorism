package main

import (
	"fmt"
	"math"
)

// 正の整数NとN個の座標値(xi, yi) (i=0,1,...,N-1)が与えられます。最も距離が近い2点間の距離を求めてください。
func main() {
	var N int
	_, err := fmt.Scan(&N)
	if err != nil {
		panic("Invalid Input")
	}
	x := make([]float64, N)
	y := make([]float64, N)
	for i := 0; i < N; i++ {
		_, err := fmt.Scan(&x[i], &y[i])
		if err != nil {
			panic("Invalid Input")
		}
	}

	minDist := 9999999999999.99

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			result := calDistance(x[i], y[i], x[j], y[j])
			if result < minDist {
				minDist = result
			}
		}
	}

	fmt.Print(minDist)
}

// calDistance は 2点間の距離を求める
func calDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}
