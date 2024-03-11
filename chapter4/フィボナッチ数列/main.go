package main

import "fmt"

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

func main() {
	var N int
	fmt.Scan(&N)
	println(fibo(N))
}
