package main

import "fmt"

//重点注意：
//1.数组的长度和类型也属于数组的一部分
//2.数组是值类型
//3.数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
//4.[n]*T表示指针数组，*[n]T表示数组指针 。

//自我总结：数组和php区别：必须定义长度和类型 其他基本一致
// func main() {
	func main(){
	/*arr 数组的三种初始化方式 start*/

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
	/*arr 数组的三种初始化方式 end*/

	/*数组的遍历 start*/
	//第一种遍历方式  索引取值
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}
	//第二种遍历方式  键值对方式
	for index, value := range arr2 {
		fmt.Println(index, value)
	}
	/*数组的遍历 end*/

	/*多维数组的定义 start*/
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println("多维数组：", a)         //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println("多维数组取值：", a[2][1]) //支持索引取值:重庆
	/*多维数组的定义 end*/

	/*多维数组的遍历 start*/
	a1 := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	for _, v1 := range a1 {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
	/*多维数组的遍历 end*/
}
