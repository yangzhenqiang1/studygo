package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	s1 := student{
		Name: "yzq",
		Age:  18,
	}
	str := "yzq"
	type1 := reflect.TypeOf(s1)
	fmt.Printf("type:%#v\n", type1)
	value1 := reflect.ValueOf(s1)
	fmt.Printf("value:%#v\n", value1)
	// k := value1.Kind()
	// fmt.Println(k)
	strT := reflect.TypeOf(str)
	strV := reflect.ValueOf(&str)
	kt := strT.Kind()
	kv := strV.Kind()
	fmt.Println("----------------------")
	fmt.Printf("%#v\n", kt)
	fmt.Printf("%#v\n", kv)

	//反射修改其原始变量需要传指针，并且使用Elem
	fmt.Println(strV)
	strV.Elem().SetString("马克图布")
	fmt.Println(str)
}
