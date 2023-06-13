package main

import (
	"fmt"
)

func main() {
	slice := []int{30, 10, 20, 60, 50, 40}
	insertSort(slice)
	fmt.Println(slice)
}

func insertSort(slice []int) {
	for i := 1; i < len(slice); i++ {
		current := slice[i]
		j := i - 1
		for ; j >= 0 && slice[j] > current; j-- {
			slice[j+1] = slice[j]
		}
		slice[j+1] = current
	}
}
