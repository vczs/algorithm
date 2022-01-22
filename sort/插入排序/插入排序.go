package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
插入排序：插入排序属于内部排序法 对要排序的元素以插入的方式找寻该元素的适当位置 以达到排序的目的。
 */

func main(){
	var slice [80000]int
	for i := 0 ; i < 80000 ; i++ {
		slice[i] = rand.Intn(100000)
	}
	startTime := time.Now().Unix()
	insertSort(&slice)
	endTime := time.Now().Unix()
	fmt.Printf("插入排序耗时：%v秒。\n",endTime-startTime)
}
func insertSort(slice *[80000]int){
	for i := 1 ; i < len(slice); i++  {
		insertVal := slice[i] //先将要排序的slice[i]赋值给insertVal
		insertIndex := i - 1 //从下标i - 1的开始比较
		//当slice[insertIndex]大于insertVal时就将slice[insertIndex]赋值给他的下一位 并将下标往左移一位再进行比较
		//直到slice[insertIndex]小于insertVal或者下标向左移动到小于0时停止前移 说明元素找到了自己的位置
		for insertIndex >= 0 && slice[insertIndex] > insertVal{
			slice[insertIndex + 1] = slice[insertIndex]
			insertIndex--
		}
		//将insertVal赋值到他适合的位置slice[insertIndex + 1]
		if insertIndex + 1 != i { //如果insertIndex + 1 != i说明元素需要移动 如果等于说明元素无需移动他就在自己合适的位置
			slice[insertIndex + 1] = insertVal
		}
	}
}
