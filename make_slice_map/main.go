package main

import "fmt"

//make 的用法  make只可以对slice map chan 使用
func main() {
	//slice
	s1 := make([]int, 1, 10)
	fmt.Println(s1)

	//map
	m1 := make(map[int]string, 1)
	fmt.Println(m1)
}
