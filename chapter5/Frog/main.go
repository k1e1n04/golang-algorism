package main

import (
	"fmt"
	"github.com/k1e1n04/golang-algorism/utils"
)

func main() {
	var N int
	fmt.Scan(&N)
	h := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&h[i])
	}

	dp := make([]int, N)
	dp[0] = 0
	for i := 1; i < N; i++ {
		if i == 1 {
			dp[1] = h[1] - h[0]
		} else {
			dp[i] = utils.Min(dp[i-1]+utils.Abs(h[i]-h[i-1]),
				dp[i-2]+utils.Abs(h[i]-h[i-2]))
		}
	}

	println(dp[N-1])
}
