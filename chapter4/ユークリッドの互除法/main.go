package main

import "fmt"

func GCD(m, n int) int {
	if n == 0 {
		return m
	}
	return GCD(n, m%n)
}

func main() {
	var M int
	var N int
	fmt.Scan(&M)
	fmt.Scan(&N)
	fmt.Println(GCD(M, N))
}
