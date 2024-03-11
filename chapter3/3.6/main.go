package main

import "fmt"

// 2の正の整数K, Nが与えられる。0<= X,Y,Z <=Kを満たす整数(X,Y,Z)の組みであって
// X + Y + Z = Nを満たすものが何通りあるかを求めるO(N^2)のアルゴリズムを設計してください。
func main() {
	var K, N int
	fmt.Scan(&K)
	fmt.Scan(&N)
	count := 0
	for i := 0; i <= K; i++ {
		for j := 0; j <= K; j++ {
			z := N - i - j
			if z >= 0 && z <= K {
				count += 1
			}
		}
	}
	println(count)
}
