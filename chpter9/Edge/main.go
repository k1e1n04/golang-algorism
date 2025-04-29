package main

import "fmt"

// Edge はグラフのエッジを表す構造体
type Edge struct {
	to     int
	weight int
}

type Graph [][]Edge

// 重み付きグラフを受け取る
func main() {
	var N, M int
	fmt.Scan(&N)
	fmt.Scan(&M)

	// ノード数Nに合わせて初期化
	g := make(Graph, N)
	for i := range g {
		g[i] = []Edge{} // 各ノードの隣接リストを初期化
	}

	for i := 0; i < M; i++ {
		var a, b, w int
		fmt.Scan(&a, &b, &w)
		// ノード番号が0始まりであることを想定
		g[a] = append(g[a], Edge{to: b, weight: w})
		g[b] = append(g[b], Edge{to: a, weight: w}) // 無向グラフならこっちも必要
	}
}
