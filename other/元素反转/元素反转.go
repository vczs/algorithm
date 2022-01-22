package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
指定数组长度，随机生成其元素并存储，再将其元素位置反转
例如：数组{0,1,2,3}，反转为数组{3,2,1,0}
*/
func main() {
	var array [6]int
	arrayLen := len(array)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < arrayLen; i++ {
		array[i] = rand.Intn(100) //随机生成0-100之内的int数据类型的数存入数组
	}
	fmt.Println("原始数组：", array)
	temp := 0
	for i := 0; i < arrayLen/2; i++ {
		temp = array[i]
		array[i] = array[arrayLen-1-i]
		array[arrayLen-1-i] = temp
	}
	fmt.Println("反转后数组：", array)
}
