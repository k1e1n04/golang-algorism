package main

import (
	"fmt"
	"github.com/k1e1n04/golang-algorism/utils"
)

func main() {
	var N int
	fmt.Scan(&N)
	a := make([]int, N)
	b := make([]int, N)
	c := make([]int, N)
	for i := range a {
		fmt.Scan(&a[i])
	}
	for i := range b {
		fmt.Scan(&b[i])
	}
	for i := range c {
		fmt.Scan(&c[i])
	}

	dp := make([][3]int, N)
	dp[0][0] = a[0]
	dp[0][1] = b[0]
	dp[0][2] = c[0]

	for i := 1; i < N; i++ {
		dp[i][0] = utils.Max(dp[i-1][1], dp[i-1][2]) + a[i]
		dp[i][1] = utils.Max(dp[i-1][0], dp[i-1][2]) + b[i]
		dp[i][2] = utils.Max(dp[i-1][0], dp[i-1][1]) + c[i]
	}

	fmt.Println(utils.Max(dp[N-1][0], utils.Max(dp[N-1][1], dp[N-1][2])))
}
