package main

import "fmt"

// main はN以下の偶数を全て列挙する
func main() {
	N := 100
	for i := 2; i <= N; i += 2 {
		fmt.Println(i)
	}
}
