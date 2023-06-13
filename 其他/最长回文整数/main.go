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
	str := "00009"

	// 将字符串中出现的数字添加到一个切片里
	s := []int{}
	// 将每个数字出现的次数存储到一个map里
	m := map[int]int{}

	for i := 0; i < len(str); i++ {
		if _, ok := m[int(str[i]-'0')]; !ok {
			s = append(s, int(str[i]-'0'))
		}
		m[int(str[i]-'0')]++
	}

	sort.Sort(sort.IntSlice(s)) // 对切片里的数字进行排序

	left := ""
	mid := ""
	right := ""

	for _, v := range s {
		if left == "" && v == 0 {
			continue
		}
		repeat := m[v]
		for i := 0; i < repeat/2; i++ {
			left = left + strconv.Itoa(v)
			right = strconv.Itoa(v) + right
		}
		if mid == "" && repeat%2 != 0 {
			mid = strconv.Itoa(v)
		}
	}

	end := left + mid + right

	fmt.Println(end)
}
