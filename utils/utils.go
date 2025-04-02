package utils

import "math"

func Min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func Abs(a int) int {
	return int(math.Abs(float64(a)))
}

func Chmin(a *int, b int) {
	if *a > b {
		*a = b
	}
}

func Chmax(a *int, b int) {
	if *a < b {
		*a = b
	}
}
