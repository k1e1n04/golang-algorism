package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	var N int
	fmt.Scan(&N)

	H := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&H[i])
	}

	S := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&S[i])
	}

	left := 0
	right := math.MaxInt

	for right-left > 1 {
		mid := (left + right) / 2

		t := make([]int, N)
		ok := true

		for i := 0; i < N; i++ {
			if mid < H[i] {
				ok = false
				break
			}
			t[i] = (mid - H[i]) / S[i] // 風船iをmid以内に割るならt[i]秒までに割る必要あり
		}

		if !ok {
			left = mid
			continue
		}

		slices.Sort(t)

		for i := 0; i < N; i++ {
			if t[i] < i { // i秒までにi番目の風船を割れない
				ok = false
				break
			}
		}

		if ok {
			right = mid
		} else {
			left = mid
		}
	}

	fmt.Println(right)
}
