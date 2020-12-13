package main

import "fmt"

func main() {

	var s1 = make([]map[int]string, 10, 10)
	fmt.Println(s1)
	s1[9] = make(map[int]string, 1)
	s1[9][1] = "a"
	fmt.Println(s1)

}
