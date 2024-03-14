package main

import "fmt"

var memo []int

func toribonacci(N int) int {
	if N == 0 {
		memo[0] = 0
		return 0
	} else if N == 1 {
		memo[1] = 0
		return 0
	} else if N == 2 {
		memo[2] = 1
		return 1
	}

	if memo[N] != -1 {
		return memo[N]
	}

	memo[N] = toribonacci(N-1) + toribonacci(N-2) + toribonacci(N-3)
	return memo[N]
}

func main() {
	var N int
	fmt.Scan(&N)
	memo = make([]int, N+1)
	for i := range memo {
		memo[i] = -1
	}
	toribonacci(N)
	fmt.Println(memo)
}
