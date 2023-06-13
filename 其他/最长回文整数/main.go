package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
给你一个仅由数字（0 - 9）组成的字符串 num 。请你找出能够使用 num 中数字形成的最大回文整数，并以字符串形式返回。该整数不含前导零 。
注意：
	无需使用 num 中的所有数字，但你必须使用 至少 一个数字。
	数字可以重新排序

示例1：
	输入：num = "444947137"
	输出："7449447"
	解释：
		从 "444947137" 中选用数字 "4449477"，可以形成回文整数 "7449447" 。
		可以证明 "7449447" 是能够形成的最大回文整数。

示例2：
	输入：num = "00009"
	输出："9"
	解释：
		可以证明 "9" 能够形成的最大回文整数。注意返回的整数不应含前导零。

*/

func main() {
	num := "00009"
	m1 := map[int]int{}

	slice := []int{}

	for i := 0; i < len(num); i++ {
		if _, ok := m1[int(num[i]-'0')]; !ok {
			slice = append(slice, int(num[i]-'0'))
		}
		m1[int(num[i]-'0')]++
	}

	sort.Sort(sort.IntSlice(slice))

	left := ""
	right := ""
	mid := ""
	for i := len(slice) - 1; i >= 0; i-- {
		k, _ := m1[slice[i]]
		for j := 0; j < k/2; j++ {
			if slice[i] == 0 && left == "" {
				continue
			}
			left += strconv.Itoa(slice[i])
			right = strconv.Itoa(slice[i]) + right
		}
		if mid == "" && k%2 != 0 {
			mid = strconv.Itoa(slice[i])
		}
	}
	end := left + mid + right
	fmt.Println(end)
}
