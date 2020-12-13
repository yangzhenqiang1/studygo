package main

import (
	"fmt"
	"unicode"
)

func main() {
	// str1 := "how do you do"
	// // s1 := strings.Field(str1)
	// s1 := strings.Fields(str1)
	// // fmt.Println(str1)
	// // fmt.Println(s1)
	// m1 := make(map[string]string, 4)
	// m2 := make(map[string]int, 4)

	// for i := 0; i < len(s1); i++ {
	// 	fmt.Println(s1[i])
	// 	for j := 0; j < len(s1); j++ {

	// 	}
	// 	value, ok := m2[s1[i]]
	// 	if ok {
	// 		m2[value] = value
	// 	}
	// }
	// fmt.Println(m2)

	//1.判断字符串中汉字的数量
	//难点判断一个汉字是汉字

	var count int //声明一个变量
	s1 := "Hellow沙河"

	//1.依次拿到字符串中字符
	for _, c := range s1 {

		//2.判断字符中是否为汉字
		if unicode.Is(unicode.Han, c) {
			count++

		}
	}

	//3.把汉字出现的次数累加起来
	fmt.Println(count)
}
