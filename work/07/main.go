package main

import "fmt"

func justifyType(x interface{}) {
	// x.(T)  断言，结合switch获得准确的类型
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func main() {

	justifyType(1)
	justifyType("1")
}
