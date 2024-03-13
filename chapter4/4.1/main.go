package main

import "fmt"

func toribonacci(N int) int {
	if N == 0 {
		return 0
	} else if N == 1 {
		return 0
	} else if N == 2 {
		return 1
	}

	return toribonacci(N-1) + toribonacci(N-2) + toribonacci(N-3)
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(toribonacci(N))
}
