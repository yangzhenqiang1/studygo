package main

import "fmt"

//arr 数组的三种初始化方式

//重点注意：数组的长度和类型也属于数组的一部分
func main() {
	//数组的第一种初始化方式
	var arr1 [30]int
	arr1 = [30]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	//第二种
	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 41, 22}
	fmt.Println(arr2)

	//第三种
	arr3 := [...]int{1: 1, 99: 122}
	fmt.Println(arr3)

}
