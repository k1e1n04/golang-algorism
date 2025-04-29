package main

import "fmt"

type Graph [][]int

func main() {
	var N, M int
	fmt.Scan(&N)
	fmt.Scan(&M)

	// ノード数Nに合わせて初期化
	g := make(Graph, N)
	for i := range g {
		g[i] = []int{} // 各ノードの隣接リストを初期化
	}

	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		// ノード番号が0始まりであることを想定
		g[a] = append(g[a], b)
		g[b] = append(g[b], a) // 無向グラフならこっちも必要
	}
}
