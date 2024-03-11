package main

import "fmt"

// N個の整数a0,a1,...aN-1と、N個の整数b0,b1...bN-1が与えられます。
// 2組の整数列からそれぞれ1個ずつ整数を選んで和をとります。
// その和として考えられる値のうち、整数K以上の範囲内で最小値を求めてください。
func main() {
	var N int
	fmt.Scan(&N)
	a := make([]int, N)
	b := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < N; i++ {
		fmt.Scan(&b[i])
	}
	var K int
	fmt.Scan(&K)
	result := 999999999999999999
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			sum := a[i] + b[j]
			if sum < result && K <= result {
				result = sum
			}
		}
	}
	fmt.Printf("anster: %s\n", result)
}
