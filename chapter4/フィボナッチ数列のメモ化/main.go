package main

import "fmt"

var memo []int

func fibo(n int) int {
	if n == 0 {
		memo[0] = 0
		return memo[0]
	} else if n == 1 {
		memo[1] = 1
		return memo[1]
	}
	if memo[n] != -1 {
		return memo[n]
	}

	memo[n] = fibo(n-1) + fibo(n-2)
	return memo[n]
}

func main() {
	var N int
	fmt.Scan(&N)
	memo = make([]int, N)
	for i := range memo {
		memo[i] = -1
	}
	fibo(N - 1)
	fmt.Println(memo)
}
