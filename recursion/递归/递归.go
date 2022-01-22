package main

import "fmt"

/**
递归就是函数或方法自己调用自己 每次调用时传入不同的变量 递归有助于编程者解决复杂的问题 让代码变得简洁

原则：
1.执行一个函数时 创建一个新的受保护的独立空间（新函数栈）
2.函数的局部变量是独立的 不会相互影响  如果想使用同一个变量 就要使用引用传递 或者全局变量
3.递归必须向退出递归的条件逼近 否则就是无限递归
4.当一个函数执行完毕，或者遇到return，就会返回；遵守谁调用，就将结果返回给谁；同时当函数执行完毕或返回时，该函数本身也会被系统销毁。
 */


/**
0:没有走过的点
1:墙
2:一条通路
3:走过的点 但是走不通
 */
func drawWay(myMap *[8][7]int , i int , j int) bool{
	if myMap[6][5] == 2 { //myMap[6][5]是出口 如果走到这个点并且是通路 说明找到出口
		return true
	}else{ //如果没到出口 就还要继续找
		if myMap[i][j] == 0 {//这个点还未走过 0表示可以探测 则需要探测周围
			//开始探测周围
			myMap[i][j] = 2 //先将他更新为2并开始探测它的周围
			if drawWay(myMap,i+1,j) {//探测下方的点 返回true表示是通路 false表示走不通
				return true
			}else if drawWay(myMap,i,j+1) { //探测右边
				return true
			}else if drawWay(myMap,i-1,j) { //探测上边
				return true
			}else if drawWay(myMap,i,j-1) { //探测左边
				return true
			}else {
				//如果上下左右都走不了 表示该点走不通 将该点更新为3  并返回false
				myMap[i][j] = 3
				return false
			}
		}else if myMap[i][j] == 1{ //1表示墙 则该点不能探测  返回false
			return false //这个点不能探测 是墙
		}else{ //该点既不是1也不是0 那就是未知的点  未知情况返回false
			fmt.Printf("未知该点：%v\n",myMap[i][j])
			return false //这个点未知
		}
	}
}
func drawMap(myMap *[8][7]int){
	//构建迷宫地图
	for i:= 0; i < len(myMap) ; i++  {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}
	myMap[3][1] = 1
	myMap[3][2] = 1
	myMap[4][3] = 1
	for i := 0 ; i < len(myMap[0]); i++ {
		myMap[0][i] = 1
		myMap[len(myMap[0])][i] = 1
	}
	fmt.Println("开始前：")
	for i := 0 ; i < len(myMap) ; i++ {
		for j := 0 ; j < len(myMap[i]) ; j++{
			fmt.Print(myMap[i][j],"  ")
		}
		fmt.Println()
	}
}
func main(){
	var myMap [8][7]int
	drawMap(&myMap)
	drawWay(&myMap,1,1) //从myMap[1][1]开始走
	fmt.Println("走完迷宫后：")
	for i := 0 ; i < len(myMap) ; i++ {
		for j := 0 ; j < len(myMap[i]) ; j++{
			fmt.Print(myMap[i][j],"  ")
		}
		fmt.Println()
	}
}