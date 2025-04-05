package main

import (
	"fmt"
	"strings"
)

func main() {
	left := 20
	right := 36

	for right-left > 1 {
		mid := left + (right-left)/2
		var answer string
		fmt.Printf("Is the age less than %d? (yes/no): ", mid)
		fmt.Scan(&answer)
		answer = strings.ToLower(strings.TrimSpace(answer))

		if answer == "yes" {
			right = mid
		} else {
			left = mid
		}
	}

	fmt.Printf("The answer is %d\n", left)
}
