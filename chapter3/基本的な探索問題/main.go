package main

import "fmt"

// N個の整数a0,a1,...aN-1と整数値vが与えられる。ai=vとなるデータが存在するかどうか判定する。
func main() {
	var N int
	fmt.Scan(&N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}
	var v int
	fmt.Scan(&v)
	result := false
	for i := 0; i < N; i++ {
		if a[i] == v {
			result = true
			break
		}
	}
	println(result)
}
