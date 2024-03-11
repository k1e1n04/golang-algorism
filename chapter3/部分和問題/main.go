package main

import "fmt"

// N個の正の整数a0,a1,...aN-1正の整数Wが与えられる。
// a0, a1,...aN-1の中から何個かの整数を選んで総和をWとすることができるかどうかを判定してください。
func main() {
	var N int
	fmt.Scan(&N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}
	var W int
	fmt.Scan(&W)

	exist := false
	for bit := 0; bit < (1 << N); bit++ {
		sum := 0
		for i := 0; i < N; i++ {
			// i 番目の要素a[i]が部分集合に含まれているかどうか
			if (bit & (1 << i)) != 0 {
				sum += a[i]
			}
		}
		if sum == W {
			exist = true
		}
	}

	if exist {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
