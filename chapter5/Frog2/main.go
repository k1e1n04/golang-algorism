package main

import (
	"cmp"
	"fmt"
	"github.com/k1e1n04/golang-algorism/utils"
	"math"
)

// グラフ上で頂点 u から 頂点 v へと遷移する辺があって
// コストをc としたときに
// chmin(&dp[v], dp[u] + c)
// とする処理を緩和と呼ぶ
func chmin[T cmp.Ordered](a *T, b T) {
	if *a > b {
		*a = b
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	h := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&h[i])
	}

	dp := make([]int, N)
	for i := 0; i < N; i++ {
		dp[i] = math.MaxInt
	}

	dp[0] = 0
	for i := 1; i < N; i++ {
		chmin(&dp[i], dp[i-1]+utils.Abs(h[i]-h[i-1]))
		if i > 1 {
			chmin(&dp[i], dp[i-2]+utils.Abs(h[i]-h[i-2]))
		}
	}

	fmt.Println(dp[N-1])
}
