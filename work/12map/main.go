package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Go Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."
	// 存储单词词频
	count := make(map[string]int)

	// 分词
	substr := strings.Split(str, " ")
	fmt.Println(substr)
	//统计词频
	for _, value := range substr {
		count[value]++
	}

	for key, value := range count {
		fmt.Printf("%s: %d \n", key, value)
	}
}
