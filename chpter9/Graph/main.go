package main

import "fmt"

type Graph [][]int

func main() {
	var N, M int
	fmt.Scan(&N)
	fmt.Scan(&M)

	g := make([][]int, N)

	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scan(&a)
		fmt.Scan(&b)

		g[a] = append(g[a], b)
		// 無向グラフの場合
		g[b] = append(g[b], a)
	}
}
