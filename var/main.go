package main

import "fmt"

var sex string //变量可以在方法体外定义，这样就不需要必须使用该变量了

//变量
//Go语言中的每一个变量都有自己的类型，并且变量必须经过声明才能开始使用。
//Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。 并且Go语言的变量声明后必须使用。
func main() {
	/*Go语言的变量声明格式为：
	var 变量名 变量类型*/

	var name string
	name = "string"
	fmt.Println(name) //如果定义后没有使用，就会报错

	/*短变量声明
	简单定义方法	根据值倒推类型 int的话 会根据当前计算机环境设置类型 例int32 int64*/
	age := 18
	fmt.Println(age)
	fmt.Printf("%T\n", age) //获取变量类型

	/*
		批量定义
		var (
		    a string
		    b int
		    c bool
		    d float32
		)
	*/

	/*在一行中 初始化定义多个变量*/
	// var title, content = "标题", "内容"

	/*声明了pi和e这两个常量之后，在整个程序运行期间它们的值都不能再发生变化了。

	多个常量也可以一起声明：

	const (
	    pi = 3.1415
	    e = 2.7182
	)
	const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。 例如：

	const (
	    n1 = 100
	    n2
	    n3
	)
	上面示例中，常量n1、n2、n3的值都是100。*/
}
