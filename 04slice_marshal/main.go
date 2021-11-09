package main

import (
	"encoding/json"
	"fmt"
)

type login struct {
	User     string `json:"user" form:"user" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func main() {
	slice1 := []string{"1", "2", "3"}
	s, err := json.Marshal(slice1)
	fmt.Printf("%#v\n", slice1)
	fmt.Printf("%#v\n", string(s))
	fmt.Printf("%#v\n", err)

	map1 := map[string]string{"User": "a", "Password": "b"}
	m1, err := json.Marshal(map1)
	fmt.Printf("%#v\n", slice1)
	fmt.Printf("%#v\n", string(m1))
	fmt.Printf("%#v\n", err)

	login1 := new(login)
	data := []byte(m1)
	err = json.Unmarshal(data, &login1)
	if err != nil {
		fmt.Printf("%#v\n", err)
	}
	fmt.Printf("%#v\n", login1)
}
