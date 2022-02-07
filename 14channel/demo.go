package main

import (
	"fmt"
	"time"
)

func ggg(aa chan int) {
	fmt.Println("aa", <-aa)

}

func main() {
	a := make(chan int)
	go ggg(a)
	time.Sleep(time.Second)
	// a <- 10
	select {
	case a <- 10:
		{
			fmt.Println("ok")
		}
	default:
		{
			fmt.Println("default")
		}
	}
	time.Sleep(time.Second)
	fmt.Println("sdsds")
}
