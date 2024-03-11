package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 各桁の値が１以上９以下の数値のみである整数とみなせるような，長さ N の文字列 S が与えられます．この文字列の中で，
// 文字と文字の間のうちのいくつかの場所に「＋」を入れることができます．１つも入れなくてもかまいませんが，「＋」が連続してはいけません．
// このようにしてできるすべての文字列を数値とみなして，総和を計算する O(2N) のアルゴリズムを設計してください．
// たとえば S＝"125" のときは，125, 1＋25 (＝26), 12＋5 (＝17), 1＋2＋5 (＝8) の総和をとって176となります．
func main() {
	var S string
	fmt.Scan(&S)
	N := len(S)
	answer := 0
	s := strings.Split(S, "")
	for bit := 0; bit < (1 << (N - 1)); bit++ {
		sum := 0
		for i := 0; i < (N - 1); i++ {
			sum *= 10
			num, _ := strconv.Atoi(s[i])
			sum += num
			if bit&(1<<i) != 0 {
				answer += sum
				sum = 0
			}
		}
		sum *= 10
		num, _ := strconv.Atoi(s[N-1])
		sum += num
		answer += sum
	}
	fmt.Println(answer)
}
