package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var W int
	fmt.Scan(&W)
	a := make([]int, N)
	for i := range a {
		fmt.Scan(&a[i])
	}

	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, W)
	}
	for i := 1; i < N; i++ {
		for j := 0; j <= W; j++ {
			
		}
	}
}
