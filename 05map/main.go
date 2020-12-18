package main

import "fmt"

func main() {
	//最常用的使用make 初始化一个map[int]string类型的参数	重点：make会分配空间
	m1 := make(map[int]string, 0)
	fmt.Println(m1 == nil)
	m1[1] = "qqq"
	m1[2] = "qqq"
	m1[3] = "qqq"
	fmt.Printf("%T\t%d\n", m1, len(m1)) //map只可以使用len()获取长度，不可以使用cao()获取容量
	fmt.Println(m1)

	fmt.Println(m1[4] == "") //获取map中没有的值，会返回一个空字符串
	//使用直接赋值声明
	m2 := map[int]int{1: 1, 3: 2, 4: 2}
	fmt.Println(m2)

	//第三种var type形式
	var m3 map[int]int
	fmt.Println(m3)

	//--------------------------------------------------------------------------------
	//取值+判断
	value, ok := m1[1]     //值存在
	fmt.Println(value, ok) //返回具体的值和一个bool类型的返回值（true）
	//取值+判断
	value, ok = m1[4]      //不存在
	fmt.Println(value, ok) //返回空字符串和一个bool类型的返回时（false）

	//--------------------------------------------------------------------------------
	//遍历
	for i := 0; i <= len(m1); i++ {
		fmt.Println(m1[i])
	}

	for k, v := range m1 { //可以使用_来省去k或者v
		fmt.Println(k, v) //预存入的顺序无关
	}

	//--------------------------------------------------------------------------------
	//map类型的切片
	m4 := map[int][]int{1: []int{1, 2, 3, 4, 5}}
	fmt.Println(m4)

	//切片类型的map
	s1 := make([]map[int]int, 10, 10)
	s1[9] = make(map[int]int, 10)
	s1[9][1] = 1
	fmt.Println(s1)

	//delete 删除map中的元素
	fmt.Println(m1)
	delete(m1, 3)
	fmt.Printf("%T\t%v", m1, m1)
}
