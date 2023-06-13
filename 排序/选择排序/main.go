package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 30, 20, 60, 50, 40}
	selectSort(slice)
	fmt.Println(slice)
}

func selectSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		max := i // 记录最大值元素的下标
		// 查找当前元素后面最大值的元素
		for j := i + 1; j < len(slice); j++ {
			if slice[max] < slice[j] {
				max = j
			}
		}
		// 判断当前元素有没有最大值 有就更换位置
		if max != i {
			slice[i], slice[max] = slice[max], slice[i]
		}
	}
}
