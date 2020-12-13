package main

import "fmt"

//make 的用法
func main() {
	//arr
	a1 := [...]int{1, 2, 3}
	fmt.Println(a1)

	//slice
	s1 := make([]int, 1, 10)
	fmt.Println(s1)

	//map
	m1 := make(map[int]string, 1)
	fmt.Println(m1)
}
