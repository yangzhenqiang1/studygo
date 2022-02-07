package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine
var wg sync.WaitGroup //定义一个全局的wg 等待组

func mySum(x, y int) {
	time := time.Now()
	sum := x + y
	fmt.Printf("%v sum:%d\n", time, sum)
	wg.Done()
}

func main() {
	a := time.Now()
	fmt.Printf("%v start\n", a)
	//通道定义方式
	var ch1 chan int          //(1)	只是声明还没有开辟空间内存 nil
	ch2 := make(chan int, 10) //(2)	声明且分配了10个缓存区的内存
	var ch12 chan int         // 声明一个传递整型的通道
	var ch22 chan bool        // 声明一个传递布尔型的通道
	var ch32 chan []int       // 声明一个传递int切片的通道-
	fmt.Println(ch1)
	fmt.Println(ch2)
	// for i := 0; i < 10; i++ { //work_pool goroutine池 开启10个goroutine
	// 	wg.Add(1)
	// 	go mySum(i, i+1) //采用goroutine方式执行
	// 	// mySum(i, i+1)
	// }
	b := time.Now()
	fmt.Printf("%v end\n", b)
	wg.Wait()
}
