package main

const N = 8

var a = []int{3, 5, 8, 10, 14, 17, 21, 39}

func binarySearch(key int) int {
	left := 0
	right := len(a) - 1
	for right >= left {
		mid := (left + (right - left)) / 2
		if a[mid] == key {
			return mid
		} else if a[mid] > key {
			right = mid - 1
		} else if a[mid] < key {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	println(binarySearch(1))
	println(binarySearch(3))
}
