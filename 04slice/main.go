package main

import "fmt"

//切片的定义	var name []T	其中name为切片的变量名，T代表切片的类型 只能存储该类型的内容
//注意：切片是引用类型
/*
切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。
就像一个圈，把底层函数给圈了起来，可以向右扩建。
*/
func main() {

	var slice1 []string        //声明一个切片	声明后不可以马上使用 切记需要初始化
	fmt.Println(slice1 == nil) //true
	slice1 = []string{}        //初始化 初始化后才可以进行赋值使用
	fmt.Println(slice1 == nil) //false
	// slice1[1] = "产生"	   //这个用法是错误的
	fmt.Println(slice1) //输出 []
	slice1 = []string{"1", "2"}
	fmt.Printf("类型：%T\t 长度：%d\t 容量：%d\n", slice1, len(slice1), cap(slice1))
	fmt.Println(slice1)
	slice1[1] = "333" //赋值时 索引不可以超出定义时的索引最大值 超出会报错 索引越界
	fmt.Println(slice1)

	arr1 := [...]int{1, 2, 3}
	fmt.Printf("类型：%T\t 长度：%d\t 容量111：%d\n", arr1, len(arr1), cap(arr1))
	fmt.Println(arr1)
	arr1[2] = 4 //其中索引（2）超出定义时的长度就会报错 索引越界
	fmt.Printf("类型：%T\t 长度：%d\t 容量：%d\n", arr1, len(arr1), cap(arr1))
	fmt.Println(arr1)

	fmt.Println("----------------------")

	/*切片还可以从数组中圈出来*/
	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 99}
	fmt.Printf("%T\t %v\t len:%d\t cap:%d\n", arr2, arr2, len(arr2), cap(arr2))
	s2 := arr2[:]  //左闭右开区间[）
	s3 := arr2[1:] //从索引为1取到最后
	// s3[7] = 10 //索引越界
	fmt.Printf("%T\t %v\t len:%d\t cap:%d\n", s2, s2, len(s2), cap(s2))
	fmt.Printf("%T\t %v\t len:%d\t cap:%d\n", s3, s3, len(s3), cap(s3))
	/*切片的本质是引用类型，像是一个圈把底层函数给圈了起来*/

	/*区别于以上的基于数组创建的切片，如果需要动态创建一个切片，可以使用内置的 make() 函数创建*/
	slice3 := make([]int, 1, 4) //slicec3代表新创建的切片的变量名称 []int 代表int类型的切片 1代表切片的长度 2代表切片的容量
	fmt.Printf("%T,%v,%d,%d", slice3, slice3, len(slice3), cap(slice3))
	// slice3[2] = 3 //会报错，因为2超出了定时时len为1的长度
	fmt.Printf("%T,%v,%d,%d", slice3, slice3, len(slice3), cap(slice3))
	/*判断切片是否为空
	要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。*/
	// fmt.Println(len(s3) == 0) //只有刚刚定义没有初始化的切片才==nil

	/*var s1 []int         //len(s1)=0;cap(s1)=0;s1==nil
	s2 := []int{}        //len(s2)=0;cap(s2)=0;s2!=nil
	s3 := make([]int, 0) //len(s3)=0;cap(s3)=0;s3!=nil*/

	/*append()为切片动态添加元素	可以添加一个两个多个*/
	slice4 := make([]int, 10)
	fmt.Println(slice4)

	/*切片和map混合使用 start*/
	// var s1 = make([]map[int]string, 10, 10)
	// fmt.Println(s1)
	// s1[9] = make(map[int]string, 1)
	// s1[9][1] = "a"
	// fmt.Println(s1)
	/*切片和map混合使用 end*/

}
