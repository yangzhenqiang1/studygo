package main

import "fmt"

func f() {
	fmt.Println("函数也可以没有返回值")
}

//常规用法 逐个定义参数 定义返回值类型
func f1(x int, y int) int {
	return x + y
}

// 连续多个类型一致的参数，可以省略非最后一个的参数类型，统一写在最后一个参数后面
func f2(x, y int) int {
	return x + y
}

//定义好返回值的变量名和类型，在函数体中可以直接使用该返回值的变量名
//并且return时可以省略该变量名，直接return
func f3(x int, y int) (ret int) {
	ret = x + y
	return
}

//可变长参数 y
//可变长参数必须放在参数的最后一个的位置
func f4(x int, y ...int) int {
	fmt.Println(y)
	fmt.Printf("%T", y) //参数中 ...int 代表是一个int类型的slice 切片
	return 1
}

//函数可以没有参数，但是有返回值
func f5() int {
	return 1
}

//函数可以返回多个值
func f6(x, y int) (m int, n int) {
	return m, n
}

func main() {

	fmt.Println("111")
	f()

	ret1 := f1(1, 2)
	fmt.Println("ret1:", ret1)

	ret2 := f2(1, 2)
	fmt.Println("ret2:", ret2)

	ret3 := f3(1, 2)
	fmt.Println("ret3:", ret3)

	ret4 := f4(1, 2, 3, 4)
	fmt.Println(ret4)
}
