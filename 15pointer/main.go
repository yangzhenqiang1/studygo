package main

import "fmt"

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	/*//指针和值的区别
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100*/

	/*//new(Type)的用法
	a := new(int)            //0
	b := new(bool)           //false
	c := new(map[int]string) //map[]
	d := new([]int)          //[]
	e := new([2]int)         //[0 0]
	f := new(chan int)       //nil
	//*a = 10
	//*b = false
	fmt.Printf("a type:%T\na value:%v\n", a, *a)
	fmt.Printf("b type:%T\nb value:%v\n", b, *b)
	fmt.Printf("c type:%T\nc value:%v\n", c, *c)
	fmt.Printf("d type:%T\nd value:%v\n", d, *d)
	fmt.Printf("e type:%T\ne value:%v\n", e, *e)
	fmt.Printf("f type:%T\nf value:%v\n", f, *f)*/

	//make 的用法
	a := make(map[int]int, 10) //初始化slice map channel
	a[1] = 1
	a[12] = 13
	fmt.Printf("a type:%T\na value:%v\n", a, a)
	s := make([]int, 1)
	fmt.Printf("s type:%T\ns value:%v\n", s, s)
	c := make(chan int, 2)
	c <- 1
	fmt.Printf("c type:%T\nc value:%v\n", c, <-c)
	//new 和 make 的区别
	//1.二者都是用来做内存分配的。
	//2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	//3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
}
