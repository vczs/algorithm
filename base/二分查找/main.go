package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(Search(slice, 100))
}

func Search(slice []int, target int) int {
	left := 0
	right := len(slice) - 1

	if len(slice) == 0 {
		return -1
	}

	for left <= right {
		mid := (left + right) / 2
		if slice[mid] == target {
			return mid
		}
		if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
