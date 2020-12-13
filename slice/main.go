package main

import "fmt"

func main() {
	// fmt.Println("环境测试")  //is ok

	// s1 := []string{}
	// s2 := make([]string, 10)
	// fmt.Println(s1)
	// fmt.Println(s2)

	// str1 := "你好我是胡妍琴"
	// fmt.Printf("%T",len(str1))

	var s1 = make([]map[int]string, 10, 10)
	fmt.Println(s1)
	s1[9] = make(map[int]string, 1)
	s1[9][1] = "a"
	fmt.Println(s1)

}
