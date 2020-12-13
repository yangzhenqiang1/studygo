package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "how do you do"
	// s1 := strings.Field(str1)
	s1 := strings.Fields(str1)
	// fmt.Println(str1)
	// fmt.Println(s1)
	m1 := make(map[string]string, 4)
	m2 := make(map[string]int, 4)

	for i := 0; i < len(s1); i++ {
		fmt.Println(s1[i])
		for j := 0; j < len(s1); j++ {
			
		}
		value, ok := m2[s1[i]]
		if ok {
			m2[value] = value
		}
	}
	fmt.Println(m2)
}
