package main

import "fmt"
/**
二分查找：二分查找法就是实现在一组有序的数字数组集合中最快找到指定元素的下标
思路：
1.先找到中间的下标midIndex := (leftIndex + rightIndex)/2，然后让中间的下标值和FindVal比较
a:如果array[midIndex] > number,那么就向leftIndex ~ (midIndex - 1)区间找
b:如果array[midIndex] < number,那么就向(midIndex + 1) ~ rightIndex区间找
2.递归执行1，直到array[midIndex] == number，则表示找到目标数的位置，那么返回其下标
3.如果LeftIndex > RightIndex，则表示找不到目标数，退出
****注意：二分查找只适用于有序数组集合****
 */
func main(){
	array := [8]int{10,20,30,40,50,60,70,80}
	binarySearch(array,0,len(array),50)
}
//四个参数：1:被查找的数组  2:左边下标  3:右边下标  4:要查找的数
func binarySearch(array [8]int, leftIndex int , rightIndex int , number int){
	if leftIndex > rightIndex {
		fmt.Println("找不到目标数") //如果LeftIndex > RightIndex，则表示找不到目标数，退出
		return
	}
	midIndex := (leftIndex + rightIndex) / 2 //先找到中间数的下标
	if number > array[midIndex] {
		binarySearch(array, midIndex+1, rightIndex, number) //如果目标数大于中间数,则递归从下标为(midIndex+1)到下标为rightIndex的数的区间找
	}else if number < array[midIndex] {
		binarySearch(array, leftIndex, midIndex-1, number) //如果目标数小于中间数,则递归从下标为leftIndex到下标为(midIndex-1)的数的区间找
	}else {
		fmt.Println(midIndex) //既不小于也不大与，那么array[midIndex] == number，则表示找到目标数的位置，那么返回其下标
	}
}