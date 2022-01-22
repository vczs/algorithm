package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
选择排序：选择排序也属于内部排序法 是从要排序的数据中 按指定规则选出某一个元素 经过和其它元素重整 再依原则交换位置后达到排序的目的
原理：
将N个数进行排序 先从下标为0-N中选最出最大的放在第0位 再从下标为1-N中选出最大的放在第1位 再从下标位2-N中选出最大的放在第二位
一共进行N-1次(前面的N-1个数比较完 最后剩下的一个肯定是最小的无需再进行比较)
上面是从大到小进行排序  从小到大排序的流程和上面的相同 只需每次挑出最小的放在进行比较的数中首位即可
*/

func main() {
	var slice [80000]int
	for i := 0; i < 80000; i++ {
		slice[i] = rand.Intn(100000)
	}
	startTime := time.Now().Unix()
	selectSort(&slice)
	endTime := time.Now().Unix()
	fmt.Printf("选择排序耗时：%v秒。\n", endTime-startTime)
}
func selectSort(slice *[80000]int) {
	for i := 0; i < len(slice)-1; i++ {
		max := slice[i]
		maxIndex := i
		for j := i; j < len(slice); j++ {
			if slice[j] < max {
				max = slice[j]
				maxIndex = j
			}
		}
		if maxIndex != i {
			slice[i], slice[maxIndex] = slice[maxIndex], slice[i]
		}
	}
}
