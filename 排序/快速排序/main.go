package main

import "fmt"

func main() {
	slice := []int{30, 10, 60, 20, 50, 40}
	quickSort(slice, 0, len(slice)-1)
	fmt.Println(slice)
}

func quickSort(slice []int, low, high int) {
	if low >= high {
		return
	}
	mid := partition(slice, low, high)
	quickSort(slice, low, mid-1)
	quickSort(slice, mid+1, high)
}

func partition(slice []int, low, high int) int {
	i := low - 1
	mid := slice[high]

	for j := low; j < high; j++ {
		if mid > slice[j] {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	// mid不用参与排序了 他的位置已明确 把他换到最中间
	slice[i+1], slice[high] = slice[high], slice[i+1]

	return i + 1
}
