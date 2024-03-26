package main

import "fmt"

// funcSolution は再帰的に問題を解くための関数
// N: 入力
// cur: 現在の値
// use: 7, 5, 3 のうちどれを使ったか
// counter: 答え
func funcSolution(N, cur int64, use byte, counter int64) int64 {
	if cur > N {
		return counter
	}

	if use == 0b111 {
		counter++
	}
	// 7 を使う
	counter = funcSolution(N, cur*10+7, use|0b001, counter)
	// 5 を使う
	counter = funcSolution(N, cur*10+5, use|0b010, counter)
	// 3 を使う
	counter = funcSolution(N, cur*10+3, use|0b100, counter)

	return counter
}

func main() {
	var N int64
	fmt.Scan(&N)

	fmt.Println(funcSolution(N, 0, 0, 0))
}
