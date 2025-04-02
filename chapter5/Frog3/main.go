package main

import (
	"fmt"
	"github.com/k1e1n04/golang-algorism/utils"
	"math"
)

// 配る遷移方式
func main() {
	var N int
	fmt.Scan(&N)
	h := make([]int, N)
	for i := range h {
		fmt.Scan(&h[i])
	}

	dp := make([]int, N)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	for i := 0; i < N; i++ {
		if i+1 < N {
			utils.Chmin(&dp[i+1], dp[i]+utils.Abs(h[i]-h[i+1]))
		}
		if i+2 < N {
			utils.Chmin(&dp[i+2], dp[i]+utils.Abs(h[i]-h[i+2]))
		}
	}

	println(dp[N-1])
}
