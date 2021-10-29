package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
快速排序：
对冒泡排序的一种改进，将要排序的数据分割为两部分，其中一部分的所有数据比零一部分的都要小
然后再按照此方法分别将两组数据再进行快速排序(分割)，整个过程递归进行直到整个数组变成有序序列
*/

func main() {
	var slice [80000]int
	for i := 0; i < 80000; i++ {slice[i] = rand.Intn(100000)}
	startTime := time.Now().Unix()
	quickSort(0, len(slice)-1, &slice)
	endTime := time.Now().Unix()
	fmt.Printf("快速排序耗时：%v秒。\n", endTime-startTime)
}
func quickSort(left int, right int, slice *[80000]int) {
	if left >= right {return}
	l := left
	r := right
	p := slice[(left+right)/2] //slice中间元素的值
	for l < r {
		for slice[l] < p {l++} //从slice[l]开始循环遍历p左边的数与p进行比较 如果小于p就l++右移比较下一位 直至找到大于p的数
		for slice[r] > p {r--} //从slice[r]开始循环遍历p右边的数与p进行比较 如果大于p就r--左移比较上一位 直至找到小于p的数
		slice[l], slice[r] = slice[r], slice[l] //将找到的 p左边大于p的数 和 p右边小于p的数交换位置
		if l >= r {break} //当l大于等于r时说明p左右已经遍历比较完 退出循环
		if slice[l] == p {r--} //如果交换位置后slice[l]等于p 就将r向左移动一位
		if slice[r] == p {l++} //如果 r向左移动一位后的slice[r]等于p的话 就将l向右移动以为
	}
	if l == r {
		l++
		r-- //如果l等于r时 就将l右移以为 r左移一位
	}
	//递归对left到r区间的数进行快速排序
	if left < r {quickSort(left, r, slice)}
	//递归对right到l区间的数进行快速排序
	if right > l {quickSort(l, right, slice)}
}
