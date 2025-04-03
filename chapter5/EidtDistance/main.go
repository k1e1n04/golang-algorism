package main

import (
	"fmt"
	"github.com/k1e1n04/golang-algorism/utils"
	"math"
)

func main() {
	var S, T string
	fmt.Scan(&S)
	fmt.Scan(&T)

	dp := make([][]int, len(S)+1)
	for i := range dp {
		dp[i] = make([]int, len(T)+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}

	// dpの初期化
	dp[0][0] = 0

	for i := 0; i <= len(S); i++ {
		for j := 0; j <= len(T); j++ {
			// 変更操作
			if i > 0 && j > 0 {
				// 変更なし
				if S[i-1] == T[j-1] {
					utils.Chmin(&dp[i][j], dp[i-1][j-1])
					// 変更
				} else {
					utils.Chmin(&dp[i][j], dp[i-1][j-1]+1)
				}
			}

			// 削除
			if i > 0 {
				utils.Chmin(&dp[i][j], dp[i-1][j]+1)
			}

			// 挿入
			if j > 0 {
				utils.Chmin(&dp[i][j], dp[i][j-1]+1)
			}

		}
	}
	fmt.Println(dp[len(S)][len(T)])
}
