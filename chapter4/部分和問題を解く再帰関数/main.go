package main

import "fmt"

func isEqual(N, W int, a []int) bool {
	if N == 0 {
		if W == 0 {
			return true
		} else {
			return false
		}
	}

	// a[i-1]を選ばない場合
	if isEqual(N-1, W, a) {
		return true
	}

	// a[i-1]を選ぶ場合
	if isEqual(N-1, W-a[N-1], a) {
		return true
	}

	// どちらもfalseの場合はfalse
	return false
}

func main() {
	var N, W int
	fmt.Scan(&N)
	fmt.Scan(&W)
	a := make([]int, N)
	for i := range a {
		fmt.Scan(&a[i])
	}
	if isEqual(N, W, a) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
