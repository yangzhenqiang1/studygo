package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	// 第一种声明方式  值方式
	var s1 student
	s1.Name = "杨振强"
	s1.Score = 18
	fmt.Printf("type:%T value:%#v\n", s1, s1)

	// 第二种	值方式
	s2 := student{
		Name:  "yzq",
		Score: 19,
	}
	fmt.Printf("type:%T value:%#v\n", s2, s2)

	// 第三种	指针方式 new操作
	s3 := new(student)
	s3.Name = "马克图布"
	s3.Score = 20
	fmt.Printf("type:%T value:%#v\n", s3, s3)

	jsonStu, err := json.Marshal(s1)
	if err != nil {
		fmt.Printf("Marshal转换失败，error：%#v\n", err)
	}
	fmt.Printf("%#v\n", jsonStu)
	fmt.Printf("%#v\n", string(jsonStu))

	jsonStu3, err3 := json.Marshal(s3)
	if err3 != nil {
		fmt.Printf("Marshal转换失败，error：%#v\n", err3)
	}
	fmt.Printf("%#v\n", jsonStu3)
	fmt.Printf("%#v\n", string(jsonStu3))

	stu1 := student{}
	str := []byte(jsonStu3)
	err = json.Unmarshal(str, &stu1)
	if err != nil {
		fmt.Printf("Unmarshal转换失败，error：%#v\n", err)
	}
	fmt.Printf("type:%T,value:%v", stu1, stu1)
}
