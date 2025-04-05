package main

import (
	"fmt"
	"github.com/k1e1n04/golang-algorism/utils"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	c := make([][]int, N+1)
	for i := range c {
		c[i] = make([]int, N+1)
		for j := range c[i] {
			fmt.Scan(&c[i][j])
		}
	}

	dp := make([]int, N+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}

	dp[0] = 0

	for i := 0; i <= N; i++ {
		for j := 0; j < i; j++ {
			utils.Chmin(&dp[i], dp[j]+c[j][i])
		}
	}

	fmt.Println(dp[N])
}
