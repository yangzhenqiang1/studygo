package main

import "fmt"

type people struct {
	name, food string
}

//定义一个函数
func a(name *string) (res string) {
	// fmt.Println("这是一个方法")
	res = *name
	return
}

//定义一个方法
func (p people) eat() (end string) {
	// fmt.Println(p.name + " eat")
	fmt.Println(p.name)
	fmt.Println(p.food)
	end = p.name + " eat " + p.food
	p.food = "apple"
	return
}

func main() {
	//函数调用
	name := "your name"
	fmt.Println(a(&name))

	//方法调用 必须由接收者发起
	yzq := people{
		"yzq",
		"meat",
	}
	fmt.Println(yzq.eat())
	fmt.Println(yzq)
}
