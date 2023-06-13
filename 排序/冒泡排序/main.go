package main

import "fmt"

func main() {
	slice := []int{10, 30, 20, 60, 50, 40}
	bubbleSort(slice)
	fmt.Println(slice)
}

func bubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1; j++ {
			if slice[j] < slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}
