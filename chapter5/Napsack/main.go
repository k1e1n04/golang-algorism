package main

import (
	"fmt"
	"github.com/k1e1n04/golang-algorism/utils"
)

// 動的計画法の部分問題の作り方の基本パターン 　N 個の対象物 {0, 1, ..., N－1} に関する問題に対して，
// 最初の i 個の対象物 {0, 1, ..., i－1} に関する問題を部分問題として考えます．
func main() {
	var N int
	fmt.Scan(&N)
	w := make([]int, N)
	for i := range w {
		fmt.Scan(&w[i])
	}
	v := make([]int, N)
	for i := range v {
		fmt.Scan(&v[i])
	}
	var W int
	fmt.Scan(&W)

	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 0; i < N; i++ {
		// jはナップサックに入る残りの重さ
		for j := 0; j <= W; j++ {
			if j-w[i] >= 0 {
				utils.Chmax(&dp[i+1][j], dp[i][j-w[i]]+v[i])
			}
			utils.Chmax(&dp[i+1][j], dp[i][j])
		}
	}

	fmt.Println(dp[N][W])
}
