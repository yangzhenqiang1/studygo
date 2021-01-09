package main

import "fmt"

// 不管是什么结构体，只要声明了speak方法，就是speaker接口类型。
type speaker interface {
	speak()
}

type cat struct{}
type dog struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func da(x speaker) {
	x.speak()
}
func main() {
	c1 := cat{}
	d1 := dog{}
	c1.speak()
	d1.speak()
	da(d1)
	da(c1)
}
