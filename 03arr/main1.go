package main

import "fmt"

//数组和其他语言区别	长度和类型作为数组的一部分，所以有一点不同就不可以比较
func main() {
	/*数组定义*/
	//第一种定义方式
	a1 := [2]string{"我", "是"}
	fmt.Println(a1)

	//第二种定义方式	最常用的
	var a2 = [...]bool{true, true, true} //... 代表从值推断出长度
	fmt.Println(a2)

	//第三种定义方式	用索引定义，可以不按照索引顺序
	a3 := [4]int{0: 1, 3: 2, 1: 2}
	fmt.Println(a3)

	/*数组遍历*/
	//第一种遍历方式
	arrs := [...]string{"a", "2", "b"}
	for i := 0; i < len(arrs); i++ {
		fmt.Println(arrs[i])
	}
	//第二种遍历方式
	for i, v := range arrs {
		fmt.Printf("%T\n", arrs[i])
		fmt.Printf("%v\n", v)
		fmt.Println(i, v)
	}
	/*定义多维数组*/
	var morearr1 [3][2]int
	morearr1 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(morearr1)

	/*定义多维数组*/
	var morearr2 [3][2]int
	morearr2 = [3][2]int{
		{1, 2}, {2}, {},
	}
	fmt.Println(morearr2)

	testarr := [...]int{1, 3, 5, 7, 8}
	var sum int
	for i := 0; i < len(testarr); i++ {
		sum += testarr[i]
	}
	fmt.Println(sum)

	for i := 0; i < len(testarr); i++ {
		for j := 0; j < i; j++ {
			if testarr[i]+testarr[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}
}
