package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Date())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// 时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	// time.Unix
	ret := time.Unix(1610174196, 0)
	fmt.Println(ret)
	fmt.Println(ret.Year())
	fmt.Println(ret.Month())

	fmt.Println(now.Add(24 * time.Hour))

	a := time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Println(a)
	a = now.Format("2006-01-02")
	fmt.Println(a)
	timeobj, err := time.Parse("2006-01-02", a)
	if err != nil {
		fmt.Printf("error：%#v", err)
	}
	fmt.Println(timeobj.Unix())

	
	fmt.Println("-------------")
	fmt.Println(now.Sub(timeobj))
	fmt.Println("-------------")
}
